package service

import (
	"context"

	"github.com/pkg/errors"

	"common/errno"
	"common/orm"
	center "common/proto/center"
	pb "common/proto/member"
	"member/idl"
	"member/model"
)

//IMember 会员服务接口
type IMember interface {
	MemberRegister(ctx context.Context, username, password string, phone int64) (int64, error)
	MemberUsernameLogin(ctx context.Context, username, password string) (*pb.MemberToken, error)
	MemberPhoneLogin(ctx context.Context, phone int64) (*pb.MemberToken, error)
	MemberEdit(ctx context.Context, id int64, userMap map[string]interface{}) error
	MemberInfo(ctx context.Context, id int64) (*pb.MemberInfo, error)
}

// MemberRegister 注册用户
func (s *Service) MemberRegister(ctx context.Context, username, password string, phone int64) (int64, error) {
	exist, err := s.repo.MemberExist(ctx, username, phone)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.member] member exist")
	}
	if exist {
		return 0, errno.ErrMemberExisted
	}
	//去中心服注册
	memberID, err := s.UserRegister(ctx, username, password, phone)
	if err != nil {
		return 0, err
	}
	u := &model.MemberModel{
		PriID:    orm.PriID{ID: memberID},
		Username: username,
		Phone:    phone,
		Status:   model.MemberStatusNormal,
	}
	err = s.repo.MemberCreate(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.member] create member")
	}
	return memberID, nil
}

// MemberUsernameLogin 用户名密码登录
func (s *Service) MemberUsernameLogin(ctx context.Context, username, password string) (*pb.MemberToken, error) {
	user, err := s.UsernameLogin(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return s.buildInfo(ctx, user)
}

// MemberPhoneLogin 手机登录
func (s *Service) MemberPhoneLogin(ctx context.Context, phone int64) (*pb.MemberToken, error) {
	user, err := s.UserPhoneLogin(ctx, phone)
	if err != nil {
		return nil, err
	}

	return s.buildInfo(ctx, user)
}

// MemberEdit 修改会员信息
func (s *Service) MemberEdit(ctx context.Context, id int64, userMap map[string]interface{}) error {
	err := s.repo.MemberUpdate(ctx, id, userMap)

	if err != nil {
		return errors.Wrapf(err, "[service.member] update member by id: %d", id)
	}
	return nil
}

// MemberInfo 获取用户信息
func (s *Service) MemberInfo(ctx context.Context, id int64) (*pb.MemberInfo, error) {
	user, err := s.memberInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return idl.TransferMember(user), nil
}

//buildInfo 构建会员信息返回
func (s *Service) buildInfo(ctx context.Context, user *center.UserToken) (*pb.MemberToken, error) {
	member, err := s.repo.GetMemberByID(ctx, user.User.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.member] by id: %v", user.User.Id)
	}
	if member == nil || member.ID == 0 { // 会员不存在，注册
		u := &model.MemberModel{
			PriID:    orm.PriID{ID: user.User.Id},
			Username: user.User.Username,
			Phone:    user.User.Phone,
			Status:   model.MemberStatusNormal,
		}
		err = s.repo.MemberCreate(ctx, u)
		if err != nil {
			return nil, errors.Wrapf(err, "[service.member] create member")
		}
	}
	return &pb.MemberToken{
		Member: idl.TransferMemberToken(user.User),
		Token:  user.Token,
	},nil
}

// memberInfo 获取会员模型
func (s *Service) memberInfo(ctx context.Context, id int64) (*model.MemberModel, error) {
	member, err := s.repo.GetMemberByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.member] from repo by id: %d", id)
	}
	if member.ID == 0 {
		return nil, errno.ErrMemberNotFound
	}
	if member.Status != model.MemberStatusNormal {
		return nil, errno.ErrMemberFrozen
	}
	return member, nil
}
