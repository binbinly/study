package service

import (
	"common/orm"
	"common/app"
	"context"
	"order/conf"
	"pkg/redis"
	"testing"
)

var srv IService

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("mall_oms")
	srv = New(conf.Conf, app.NewTestClient("192.168.8.76:8500"))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_OrderSubmit(t *testing.T) {
	type args struct {
		ctx       context.Context
		memberID  int64
		addressID int64
		couponID  int64
		skuIds    []int64
		note      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"OrderSubmit",
			args: args{
				ctx:       context.Background(),
				memberID:  3,
				addressID: 2,
				couponID:  1,
				skuIds:    []int64{12900},
				note:      "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.OrderSubmit(tt.args.ctx, tt.args.memberID, tt.args.addressID, tt.args.couponID, tt.args.skuIds, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderSubmit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestService_SubmitSkuOrder(t *testing.T) {
	got, err := srv.SubmitSkuOrder(context.Background(), 3, 12900, 2, 1, 1, "")
	if err != nil {
		t.Errorf("OrderSubmit() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_OrderDetail(t *testing.T) {
	got, err := srv.OrderDetail(context.Background(), 2, 3)
	if err != nil {
		t.Errorf("OrderDetail() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_OrderCancel(t *testing.T) {
	err := srv.OrderCancel(context.Background(), 2, 3)
	if err != nil {
		t.Errorf("OrderCancel() error = %v", err)
		return
	}
}