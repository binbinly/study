package service

import (
	"chat-micro/pkg/crypt"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
)

// ReportCreate 举报好友/群
func (s *Service) ReportCreate(ctx context.Context, UserID, friendID uint32, cType int8, cat, content string) error {
	is, err := s.repo.ReportExistPending(ctx, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.report] exist id:%d", friendID)
	}
	if is { // 已举报过
		return ErrReportExisted
	}
	report := &model.ReportModel{
		UID:        orm.UID{UserID: UserID},
		TargetID:   friendID,
		TargetType: cType,
		Content:    content,
		Category:   cat,
		Status:     model.ReportStatusPending,
	}
	_, err = s.repo.ReportCreate(ctx, report)
	if err != nil {
		return errors.Wrapf(err, "[service.report] create report err")
	}
	return nil
}

//GetUploadSIgnUrl 获取资源上传地址
func (s *Service) GetUploadSIgnUrl(ctx context.Context, name string) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", errors.Wrap(err, "new uuid")
	}
	key := fmt.Sprintf("images/%s/%s%s", time.Now().Format("20060102"),
		crypt.Md5ToString(id.String()), getFileExt(name))
	return s.opts.storage.SignUrl(ctx, key)
}

// 获取文件后缀名
func getFileExt(filename string) string {
	index := strings.LastIndex(filename, ".")
	//comma的意思是从字符串tracer查找第一个逗号，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么逗号的位置就是12
	return filename[index:]
}