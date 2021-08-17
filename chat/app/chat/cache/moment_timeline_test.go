package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimelineCache_GetCache(t *testing.T) {
	cache := NewTimelineCache()
	ctx := context.Background()
	err := cache.SetCache(ctx, 1, 2, 10)
	assert.NoError(t, err)
	err = cache.SetCache(ctx, 3, 4, 20)
	assert.NoError(t, err)
	data, err := cache.GetCache(ctx, 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, data, int64(10))
}
