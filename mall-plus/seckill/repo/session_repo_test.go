package repo

import (
	"context"
	"pkg/redis"
	"testing"
)

var repo IRepo

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	repo = New(redis.Client)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_GetSessionAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"GetSessionAll",
			args: args{ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetSessionAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}
