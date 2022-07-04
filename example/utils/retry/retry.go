package retry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	// DefaultRetryTimes times of retry
	DefaultRetryTimes = 5
	// DefaultRetryDuration time duration of two retries
	DefaultRetryDuration = time.Second * 3
)

// Func is function that retry executes
type Func func() error

// Options is config for retry
type Options struct {
	context  context.Context
	times    uint
	duration time.Duration
}

// Option is for adding retry config
type Option func(options *Options)

// WIthContext set retry context config
func WIthContext(ctx context.Context) Option {
	return func(o *Options) {
		o.context = ctx
	}
}

// WithTimes set times of retry
func WithTimes(n uint) Option {
	return func(o *Options) {
		o.times = n
	}
}

// WithDuration set duration of retries
func WithDuration(d time.Duration) Option {
	return func(o *Options) {
		o.duration = d
	}
}

// Retry executes the retryFunc repeatedly until it was successful or canceled by the context
// The default times of retries is 5 and the default duration between retries is 3 seconds
func Retry(fn Func, opts ...Option) error {
	os := &Options{
		context:  context.TODO(),
		times:    DefaultRetryTimes,
		duration: DefaultRetryDuration,
	}

	for _, opt := range opts {
		opt(os)
	}

	var i uint
	for i < os.times {
		err := fn()
		if err != nil {
			select {
			case <-time.After(os.duration):
			case <-os.context.Done():
				return errors.New("retry is cancelled")
			}
		} else {
			return nil
		}
		i++
	}

	funcPath := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	lastSlash := strings.LastIndex(funcPath, "/")
	funcName := funcPath[lastSlash+1:]

	return fmt.Errorf("function %s run failed after %d times retry", funcName, i)
}
