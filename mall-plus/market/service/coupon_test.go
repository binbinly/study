package service

import (
	"context"
	"testing"
)

func TestService_GetCouponList(t *testing.T) {
	type args struct {
		ctx      context.Context
		memberID int64
		skuID    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"GetCouponList",
			args: args{
				ctx:      context.Background(),
				memberID: 3,
				skuID:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetCouponList(tt.args.ctx, tt.args.memberID, tt.args.skuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCouponList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len: %v", len(got))
		})
	}
}

func TestService_GetMyCouponList(t *testing.T) {
	got, err := srv.GetMyCouponList(context.Background(), 3)
	if err != nil {
		t.Errorf("GetMyCouponList() error = %v", err)
		return
	}
	t.Logf("got len: %v", len(got))
}

func TestService_CouponDraw(t *testing.T) {
	err := srv.CouponDraw(context.Background(), 3, 1)
	if err != nil {
		t.Errorf("CouponDraw() error = %v", err)
		return
	}
}

func TestService_GetCouponInfo(t *testing.T) {
	got, err := srv.GetCouponInfo(context.Background(), 3, 1)
	if err != nil {
		t.Errorf("GetCouponInfo() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_CouponUsed(t *testing.T) {
	err := srv.CouponUsed(context.Background(), 3, 1, 1)
	if err != nil {
		t.Errorf("CouponUsed() error = %v", err)
		return
	}
}