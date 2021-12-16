package service

import (
	"common/orm"
	"context"
	"market/conf"
	"pkg/redis"
	"testing"
)

var srv IService

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("mall_sms")
	srv = New(conf.Conf)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_GetHomeCatData(t *testing.T) {
	type args struct {
		ctx context.Context
		cid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetHomeCatData",
			args: args{
				ctx: context.Background(),
				cid: 1220,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetHomeCatData(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHomeCatData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(got))
		})
	}
}

func TestService_GetHomeData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"GetHomeData",
			args: args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetHomeData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHomeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len: %v", len(got))
		})
	}
}

func TestService_GetNoticeList(t *testing.T) {
	type args struct {
		ctx    context.Context
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetNoticeList",
			args: args{
				ctx:    context.Background(),
				offset: 0,
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetNoticeList(tt.args.ctx, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoticeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len: %v", len(got))
		})
	}
}

func TestService_GetPayConfig(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"GetPayConfig",
			args: args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetPayConfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPayConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len: %v", len(got))
		})
	}
}

func TestService_GetSearchData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSearchData",
			args: args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := srv.GetSearchData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSearchData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}
