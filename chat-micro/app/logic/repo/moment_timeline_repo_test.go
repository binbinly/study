package repo

import (
	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
	"context"
	"gorm.io/gorm"
	"testing"
)

func TestRepo_TimelineBatchCreate(t *testing.T) {
	type args struct {
		ctx    context.Context
		tx     *gorm.DB
		models []*model.MomentTimelineModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"TimelineBatchCreate",
			args: args{
				ctx:    context.Background(),
				tx:     orm.GetDB(),
				models: []*model.MomentTimelineModel{
					{
						UID:        orm.UID{UserID: 1},
						MomentID:   1,
						IsOwn:      0,
					},
					{
						UID:        orm.UID{UserID: 2},
						MomentID:   1,
						IsOwn:      0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := r.TimelineBatchCreate(tt.args.ctx, tt.args.tx, tt.args.models)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimelineBatchCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotIds len: %v", len(gotIds))
		})
	}
}

func TestRepo_TimelineExist(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   uint32
		momentID uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"TimelineExist",
			args: args{
				ctx:      context.Background(),
				userID:   1,
				momentID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIs, err := r.TimelineExist(tt.args.ctx, tt.args.userID, tt.args.momentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimelineExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotIs: %v", gotIs)
		})
	}
}