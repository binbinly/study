package repo

import (
	"chat-micro/app/logic/model"
	"chat-micro/internal/cache"
	"chat-micro/internal/orm"
	"chat-micro/pkg/redis"
	"context"
	"gorm.io/gorm"
	"testing"
)

var r IRepo

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("chat")
	r = New(orm.GetDB(), cache.NewCache())
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_ApplyCreate(t *testing.T) {
	type args struct {
		ctx   context.Context
		apply model.ApplyModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ApplyCreate",
			args: args{
				ctx: context.Background(),
				apply: model.ApplyModel{
					UID:      orm.UID{UserID: 1},
					FriendID: 2,
					Nickname: "test",
					LookMe:   0,
					LookHim:  0,
					Status:   0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.ApplyCreate(tt.args.ctx, tt.args.apply)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId:%v", gotId)
		})
	}
}

func TestRepo_ApplyPendingCount(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"ApplyPendingCount",
			args:args{
				ctx:    context.Background(),
				userID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := r.ApplyPendingCount(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyPendingCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotC:%v", gotC)
		})
	}
}

func TestRepo_ApplyUpdateStatus(t *testing.T) {
	type args struct {
		ctx      context.Context
		tx       *gorm.DB
		id       uint32
		friendID uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"ApplyUpdateStatus",
			args: args{
				ctx:      context.Background(),
				tx:       orm.GetDB(),
				id:       1,
				friendID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.ApplyUpdateStatus(tt.args.ctx, tt.args.tx, tt.args.id, tt.args.friendID); (err != nil) != tt.wantErr {
				t.Errorf("ApplyUpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetApplyByFriendID(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   uint32
		friendID uint32
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
	}{
		{
			name:"GetApplyByFriendID",
			args: args{
				ctx:      context.Background(),
				userID:   1,
				friendID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApply, err := r.GetApplyByFriendID(tt.args.ctx, tt.args.userID, tt.args.friendID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplyByFriendID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotApply:%v", gotApply)
		})
	}
}

func TestRepo_GetApplysByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint32
		offset int
		limit  int
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			name:"GetApplysByUserID",
			args: args{
				ctx:    context.Background(),
				userID: 3,
				offset: 0,
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetApplysByUserID(tt.args.ctx, tt.args.userID, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApplysByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}