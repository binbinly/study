package function

import "time"

// Watcher is used for record code execution time
type Watcher struct {
	startTime int64
	stopTime  int64
	execution  bool
}

// Start the watch timer.
func (w *Watcher) Start() {
	w.startTime = time.Now().UnixNano()
	w.execution = true
}

// Stop the watch timer.
func (w *Watcher) Stop() {
	w.stopTime = time.Now().UnixNano()
	w.execution = false
}

// GetElapsedTime get execute elapsed time.
func (w *Watcher) GetElapsedTime() time.Duration {
	if w.execution {
		return time.Duration(time.Now().UnixNano() - w.startTime)
	}
	return time.Duration(w.stopTime - w.startTime)
}

// Reset the watch timer.
func (w *Watcher) Reset() {
	w.startTime = 0
	w.stopTime = 0
	w.execution = false
}