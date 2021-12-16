package service

import (
	"common/orm"
	"common/app"
	"context"
	"member/conf"
	"pkg/redis"
	"testing"
)

var srv IService

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("mall_ums")
	srv = New(conf.Conf, app.NewTestClient("192.168.8.76:8500"))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_MemberRegister(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
		phone    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "MemberRegister",
			args: args{
				ctx:      context.Background(),
				username: "test1",
				password: "123456",
				phone:    15555557777,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := srv.MemberRegister(tt.args.ctx, tt.args.username, tt.args.password, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId: %v", gotId)
		})
	}
}

func TestService_MemberUsernameLogin(t *testing.T) {
	got, _, err := srv.MemberUsernameLogin(context.Background(), "test1", "123456")
	if err != nil {
		t.Errorf("MemberUsernameLogin() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_MemberPhoneLogin(t *testing.T) {
	got, _, err := srv.MemberPhoneLogin(context.Background(), 15555557777)
	if err != nil {
		t.Errorf("MemberPhoneLogin() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_MemberEdit(t *testing.T) {
	err := srv.MemberEdit(context.Background(), 3, map[string]interface{}{"nickname":"测试"})
	if err != nil {
		t.Errorf("MemberEdit() error = %v", err)
		return
	}
}

func TestService_MemberInfo(t *testing.T) {
	got, err := srv.MemberInfo(context.Background(), 3)
	if err != nil {
		t.Errorf("MemberInfo() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}