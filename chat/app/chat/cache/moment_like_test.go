package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikeCache_MultiGetCache(t *testing.T) {
	cache := NewLikeCache()
	ctx := context.Background()
	err := cache.SetCache(ctx, 1, []uint32{1, 2, 3})
	assert.NoError(t, err)
	err = cache.SetCache(ctx, 2, []uint32{4, 5, 6})
	assert.NoError(t, err)
	data, err := cache.GetCache(ctx, 1)
	assert.NoError(t, err)
	t.Logf("data:%v:%v", data, len(*data))
	datas, err := cache.MultiGetCache(ctx, []uint32{1, 2})
	assert.NoError(t, err)
	t.Logf("datas:%#v", datas)
}
