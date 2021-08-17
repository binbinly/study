package center

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"chat/pkg/log"
	"chat/pkg/redis"
)

const (
	_verifyCodeRedisKey = "app:vcode:%v"   // 验证码key
	_maxDurationTime    = 10 * time.Minute // 验证码有效期
)

var (
	//ErrVerifyCodeRuleMinute 发送验证码受限
	ErrVerifyCodeRuleMinute = errors.New("sms:minute limit")
	//ErrVerifyCodeRuleHour 小时限制
	ErrVerifyCodeRuleHour = errors.New("sms:hour limit")
	//ErrVerifyCodeRuleDay 天级限制
	ErrVerifyCodeRuleDay = errors.New("sms:day limit")
	//ErrVerifyCodeNotMatch 验证码不匹配
	ErrVerifyCodeNotMatch = errors.New("code:empty")
)

//限制规则
var rules = []*rule{{
	Count: 1,
	TTL:   60,
	Key:   "sms:rule_minute:",
	Err:   ErrVerifyCodeRuleMinute,
}, {
	Count: 10,
	TTL:   3600,
	Key:   "sms:rule_hour:",
	Err:   ErrVerifyCodeRuleHour,
}, {
	Count: 15,
	TTL:   86400,
	Key:   "sms:rule_day:",
	Err:   ErrVerifyCodeRuleDay,
}}

//限制规则
type rule struct {
	Count int           //限制次数
	TTL   time.Duration //限制时间
	Key   string        //key
	Err   error         //错误
}

// SendSMS 发送短信
func (c *Center) SendSMS(ctx context.Context, phone string) (string, error) {
	code, err := c.genVCode(ctx, phone)
	if err != nil {
		return "", err
	}
	if c.c.Sms.IsReal { // 调用第三方发送服务
		if err = c.checkRules(ctx, phone); err != nil {
			return "", err
		}
		err = c.realSend(phone)
		if err != nil {
			return "", err
		}
		c.execRules(ctx, phone)
		return "", nil
	}
	return code, nil
}

// CheckVCode 验证校验码是否正确
func (c *Center) CheckVCode(ctx context.Context, phone int64, vCode string) error {
	oldVCode, err := c.getVCode(ctx, phone)
	if err != nil {
		return errors.Wrapf(err, "[center.code] get verify code")
	}

	if vCode != oldVCode {
		return ErrVerifyCodeNotMatch
	}

	return nil
}

// checkRules 验证规则
func (c *Center) checkRules(ctx context.Context, phone string) (err error) {
	if !c.c.Sms.IsReal {
		return nil
	}

	var r string
	for _, v := range rules {
		r, err = redis.Client.Get(ctx, v.Key+phone).Result()
		if err == redis.Nil {
			return nil
		} else if err != nil {
			return errors.Wrap(err, "[center.code] redis get rule err")
		}
		num, err := strconv.Atoi(r)
		if err != nil {
			return errors.Wrapf(err, "[center.code] atoi r:%v", r)
		}
		if num >= v.Count {
			return v.Err
		}
	}
	return nil
}

// execRules 发送成功,执行限流规则
func (c *Center) execRules(ctx context.Context, phone string) {
	if !c.c.Sms.IsReal {
		return
	}
	var err error
	for _, v := range rules {
		pipe := redis.Client.Pipeline()
		pipe.Incr(ctx, v.Key+phone)
		pipe.Expire(ctx, v.Key+phone, v.TTL*time.Second)
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Warnf("[service.sms] redis pipe exec err:%v", err)
		}
	}
}

//realSend 真实发送短信
func (c *Center) realSend(phone string) error {
	//TODO 短信服务
	return nil
}

// genVCode 生成验证码
func (c *Center) genVCode(ctx context.Context, phone string) (string, error) {
	// step1: 生成随机数
	vCodeStr := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// step2: 写入到redis里
	// 使用set, key使用前缀+手机号 缓存10分钟）
	key := fmt.Sprintf(_verifyCodeRedisKey, phone)
	err := redis.Client.Set(ctx, key, vCodeStr, _maxDurationTime).Err()
	if err != nil {
		return "", errors.Wrap(err, "[center.code] redis set verify code err")
	}

	return vCodeStr, nil
}

// getVCode 获取验证码
func (c *Center) getVCode(ctx context.Context, phone int64) (string, error) {
	// 直接从redis里获取
	key := fmt.Sprintf(_verifyCodeRedisKey, phone)
	verifyCode, err := redis.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", errors.Wrap(err, "[center.code] redis get verify code err")
	}

	return verifyCode, nil
}
