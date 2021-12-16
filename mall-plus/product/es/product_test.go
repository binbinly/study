package es

import (
	"context"
	"pkg/database/elasticsearch"
	"testing"
)

var product *Product

func TestMain(m *testing.M) {
	product = New(elasticsearch.NewClient(&elasticsearch.Config{
		Sniff:    false,
		Host:     "http://192.168.8.76:9200",
	}))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestProduct_Init(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"Init",
			args: args{ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := product.Init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_Search(t *testing.T) {
	query := NewQuery()
	query.Keyword("桌子").FilterCatID(953).
		FilterBrandIds([]int64{347,348}).
		FilterHasStock(false).
		FilterAttrs(26, []string{"艺美嘉","摩登奢居"}).
		FilterSkuPrice(10, 200000)
	res, err := product.Search(context.Background(), query, 0, 20, "skuPrice", true)
	if err != nil {
		t.Errorf("Search() error = %v", err)
	}
	t.Logf("res: %v", res)
}

func TestProduct_Query(t *testing.T) {
	res, err := product.Query(context.Background(), 0, 20)
	if err != nil {
		t.Errorf("Query() error = %v", err)
	}
	t.Logf("res: %v", res)
}