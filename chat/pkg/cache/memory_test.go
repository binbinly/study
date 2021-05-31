package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewMemoryCache(t *testing.T) {
	asserts := assert.New(t)

	client, err := NewMemoryCache("memory-unit-test", JSONEncoding{})
	asserts.NoError(err)
	asserts.NotNil(client)
}

func TestMemoStore_Set(t *testing.T) {
	asserts := assert.New(t)

	store, err := NewMemoryCache("memory-unit-test", JSONEncoding{})
	asserts.NoError(err)
	err = store.Set("test-key", "test-val", -1)
	asserts.NoError(err)
}

func TestMemoStore_Get(t *testing.T) {
	asserts := assert.New(t)
	store, err := NewMemoryCache("memory-unit-test", JSONEncoding{})
	asserts.NoError(err)
	// 正常情况
	{
		var gotVal string
		setVal := "test-val"
		err := store.Set("test-get-key", setVal, 3600)
		asserts.NoError(err)

		// wait for value to pass through buffers
		time.Sleep(10 * time.Millisecond)

		err = store.Get("test-get-key", &gotVal)
		asserts.NoError(err)
		t.Log(setVal, gotVal)
		asserts.Equal(setVal, gotVal)
	}
}