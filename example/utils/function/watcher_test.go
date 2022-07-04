package function

import (
	"testing"

	"lib/utils/internal"
)

func TestWatcher(t *testing.T) {
	assert := internal.NewAssert(t, "TestWatcher")

	w := &Watcher{}
	w.Start()

	longRunningTask()

	assert.Equal(true, w.execution)

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	t.Log("Elapsed Time (milsecond)", eapsedTime)

	assert.Equal(false, w.execution)

	w.Reset()

	assert.Equal(int64(0), w.startTime)
	assert.Equal(int64(0), w.stopTime)
}

func longRunningTask() {
	var slice []int64
	for i := 0; i < 10000000; i++ {
		slice = append(slice, int64(i))
	}
}
