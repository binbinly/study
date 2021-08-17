package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentCache_GetCache(t *testing.T) {
	cache := NewCommentCache()
	data, err := cache.MultiGetCache(context.Background(), []uint32{1, 2, 3})
	assert.NoError(t, err)
	t.Logf("data:%#v", data)
}
