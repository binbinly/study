package cache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"lib/pkg/encoding"
)

func Test_NewMemoryCache(t *testing.T) {
	asserts := assert.New(t)

	client := NewMemoryCache("memory-unit-test", encoding.JSONEncoding{})
	asserts.NotNil(client)
}

func TestMemoStore_Set(t *testing.T) {
	asserts := assert.New(t)

	store := NewMemoryCache("memory-unit-test", encoding.JSONEncoding{})
	err := store.Set(context.Background(), "test-key", "test-val", -1)
	asserts.NoError(err)
}

func TestMemoStore_Get(t *testing.T) {
	asserts := assert.New(t)
	store := NewMemoryCache("memory-unit-test", encoding.JSONEncoding{})
	ctx := context.Background()

	// 正常情况
	{
		var gotVal string
		setVal := "test-val"
		err := store.Set(ctx, "test-get-key", setVal, time.Minute)
		asserts.NoError(err)

		// wait for value to pass through buffers
		time.Sleep(10 * time.Millisecond)

		err = store.Get(ctx, "test-get-key", &gotVal)
		asserts.NoError(err)
		t.Log(setVal, gotVal)
		asserts.Equal(setVal, gotVal)
	}
}
