package service

import (
	"common/orm"
	"context"
	"pkg/redis"
	"testing"
	"warehouse/conf"
)

var srv IService

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("mall_wms")
	srv = New(conf.Conf)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_GetSkuStock(t *testing.T) {
	type args struct {
		ctx   context.Context
		skuID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSkuStock",
			args: args{
				ctx:   context.Background(),
				skuID: 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetSkuStock(tt.args.ctx, tt.args.skuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestService_GetSpuStock(t *testing.T) {
	stocks, err := srv.GetSpuStock(context.Background(), 1, []int64{13, 14, 15})
	if err != nil {
		t.Errorf("GetSpuStock() error = %v", err)
		return
	}
	t.Logf("stocks: %v", stocks)
}

func TestService_SKuStockLock(t *testing.T) {
	err := srv.SKuStockLock(context.Background(), 2, "123", "test", "test", "test", "test", map[int64]int32{13: 5, 12900: 5})
	if err != nil {
		t.Errorf("SKuStockLock() error = %v", err)
		return
	}
}

func TestService_SkuStockUnlock(t *testing.T) {
	err := srv.SkuStockUnlock(context.Background(), 2, true)
	if err != nil {
		t.Errorf("SkuStockUnlock() error = %v", err)
		return
	}
}