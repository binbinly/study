package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"chat/pkg/redis"
)

const (
	verifyCodeRedisKey = "app:vcode:%v"   // 验证码key
	maxDurationTime    = 10 * time.Minute // 验证码有效期
)

//限制规则
var rules = []*rule{{
	Count: 1,
	Ttl:   60,
	Key:   "sms:rule_minute:",
	Err:   ErrVerifyCodeRuleMinute,
}, {
	Count: 10,
	Ttl:   3600,
	Key:   "sms:rule_hour:",
	Err:   ErrVerifyCodeRuleHour,
}, {
	Count: 15,
	Ttl:   86400,
	Key:   "sms:rule_day:",
	Err:   ErrVerifyCodeRuleDay,
}}

//限制规则
type rule struct {
	Count int           //限制次数
	Ttl   time.Duration //限制时间
	Key   string        //key
	Err   error        //错误
}

// SendSMS 发送短信
func (s *Service) SendSMS(phone string) (string, error) {
	code, err := s.genVCode(phone)
	if err != nil {
		return "", err
	}
	if s.c.Sms.IsReal {// 调用第三方发送服务
		if err = s.checkRules(phone); err != nil {
			return "", err
		}
		err = s.realSend(phone)
		if err != nil {
			return "", err
		}
		s.execRules(phone)
		return "", nil
	}
	return code, nil
}

// CheckVCode 验证校验码是否正确
func (s *Service) CheckVCode(phone int64, vCode string) error {
	oldVCode, err := s.getVCode(phone)
	if err != nil {
		return errors.Wrapf(err, "[service.code] get verify code")
	}

	if vCode != oldVCode {
		return ErrVerifyCodeNotMatch
	}

	return nil
}

// checkRules 验证规则
func (s *Service) checkRules(phone string) error {
	if !s.c.Sms.IsReal {
		return nil
	}
	for _, v := range rules {
		rule, err := redis.Client.Get(v.Key + phone).Result()
		if err == redis.Nil {
			return nil
		} else if err != nil {
			return errors.Wrap(err, "[service.code] redis get rule err")
		}
		c, err := strconv.Atoi(rule)
		if err != nil {
			return errors.Wrapf(err, "[service.code] atoi err rule:%v", rule)
		}
		if c >= v.Count {
			return v.Err
		}
	}
	return nil
}

// execRules 发送成功,执行规则
func (s *Service) execRules(phone string) {
	if !s.c.Sms.IsReal {
		return
	}
	for _, v := range rules {
		redis.Client.Incr(v.Key + phone)
		redis.Client.Expire(v.Key+phone, v.Ttl*time.Second)
	}
}

//realSend 真实发送短信
func (s *Service) realSend(phone string) error {
	//TODO 短信服务
	return nil
}

// genVCode 生成验证码
func (s *Service) genVCode(phone string) (string, error) {
	// step1: 生成随机数
	vCodeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// step2: 写入到redis里
	// 使用set, key使用前缀+手机号 缓存10分钟）
	key := fmt.Sprintf(verifyCodeRedisKey, phone)
	err := redis.Client.Set(key, vCodeStr, maxDurationTime).Err()
	if err != nil {
		return "", errors.Wrap(err, "[service.code] redis set verify code err")
	}

	return vCodeStr, nil
}

// getVCode 获取验证码
func (s *Service) getVCode(phone int64) (string, error) {
	// 直接从redis里获取
	key := fmt.Sprintf(verifyCodeRedisKey, phone)
	verifyCode, err := redis.Client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", errors.Wrap(err, "[service.code] redis get verify code err")
	}

	return verifyCode, nil
}
