package service

import (
	"cart/conf"
	"common/app"
	"context"
	"pkg/redis"
	"testing"
)

var srv IService

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	srv = New(conf.Conf, app.NewTestClient("192.168.8.76:8500"))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_AddCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int64
		skuID  int64
		num    int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"AddCart",
			args: args{
				ctx:    context.Background(),
				userID: 3,
				skuID:  12900,
				num:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := srv.AddCart(tt.args.ctx, tt.args.userID, tt.args.skuID, tt.args.num); (err != nil) != tt.wantErr {
				t.Errorf("AddCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_EditCart(t *testing.T) {
	err := srv.EditCart(context.Background(), 3, 12900, 12901, 2)
	if err != nil {
		t.Errorf("EditCart() error = %v", err)
	}
}

func TestService_EditCartNum(t *testing.T) {
	err := srv.EditCartNum(context.Background(), 3, 12900, 2)
	if err != nil {
		t.Errorf("EditCartNum() error = %v", err)
	}
}

func TestService_CartList(t *testing.T) {
	got, err := srv.CartList(context.Background(), 3)
	if err != nil {
		t.Errorf("CartList() error = %v", err)
	}
	t.Logf("got len: %v", len(got))
}

func TestService_BatchGetCarts(t *testing.T) {
	got, err := srv.BatchGetCarts(context.Background(), 3, []int64{12900, 12902})
	if err != nil {
		t.Errorf("BatchGetCarts() error = %v", err)
	}
	t.Logf("got len: %v", len(got))
}

func TestService_DelCart(t *testing.T) {
	err := srv.DelCart(context.Background(), 3, []int64{12900, 12902})
	if err != nil {
		t.Errorf("DelCart() error = %v", err)
	}
}

func TestService_ClearCart(t *testing.T) {
	err := srv.ClearCart(context.Background(), 3)
	if err != nil {
		t.Errorf("ClearCart() error = %v", err)
	}
}