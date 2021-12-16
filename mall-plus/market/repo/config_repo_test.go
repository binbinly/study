package repo

import (
	"common/orm"
	"common/util"
	"context"
	"market/model"
	"pkg/redis"
	"testing"
)

var repo IRepo

func TestMain(m *testing.M) {
	redis.InitTestRedis()
	orm.InitTest("mall_sms")
	repo = New(orm.GetDB(), util.NewCache())
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_GetConfigByName(t *testing.T) {
	type fields struct {
		Repo util.Repo
	}
	type args struct {
		ctx  context.Context
		name string
		v    interface{}
	}
	fs := fields{Repo: util.Repo{
		DB:    orm.GetDB(),
		Cache: util.NewCache(),
	}}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: model.ConfigKeyHomeCat,
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: model.ConfigKeyHomeCat,
				v:    []*model.ConfigHomeCat{},
			},
			wantErr: false,
		},
		{
			name: model.ConfigKeyPayList,
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: model.ConfigKeyPayList,
				v:    []*model.ConfigPayList{},
			},
			wantErr: false,
		},
		{
			name: "test",
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: "test",
				v:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				Repo: tt.fields.Repo,
			}
			if err := r.GetConfigByName(tt.args.ctx, tt.args.name, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("GetConfigByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
