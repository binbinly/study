package repo

import (
	"chat-micro/app/logic/model"
	"context"
	"testing"
)

func TestRepo_GetTagNamesByIds(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint32
		ids    []uint32
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
	}{
		{
			name:"GetTagNamesByIds",
			args: args{
				ctx:    context.Background(),
				userID: 1,
				ids:    []uint32{1,2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames, err := r.GetTagNamesByIds(tt.args.ctx, tt.args.userID, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagNamesByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotNames: %v", gotNames)
		})
	}
}

func TestRepo_GetTagsByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uint32
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			name:"GetTagsByUserID",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetTagsByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_TagBatchCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		tags []*model.UserTagModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"TagBatchCreate",
			args: args{
				ctx:  context.Background(),
				tags: []*model.UserTagModel{
					{
						UserID: 1,
						Name:   "aaa",
					},
					{
						UserID: 2,
						Name:   "bb",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := r.TagBatchCreate(tt.args.ctx, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagBatchCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotIds: %v", gotIds)
		})
	}
}
