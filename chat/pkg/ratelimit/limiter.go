package ratelimit

import (
	"context"
)

// Op operations type.
type Op int

const (
	// Success operation type: success
	Success Op = iota
	// Ignore operation type: ignore
	Ignore
	// Drop operation type: drop
	Drop
)

//AllowOptions 结构
type AllowOptions struct{}

//AllowOption allow options.
type AllowOption interface {
	Apply(*AllowOptions)
}

// DoneInfo done info.
type DoneInfo struct {
	Err error
	Op  Op
}

// DefaultAllowOpts returns the default allow options.
func DefaultAllowOpts() AllowOptions {
	return AllowOptions{}
}

// Limiter limit interface.
type Limiter interface {
	Allow(ctx context.Context, opts ...AllowOption) (func(info DoneInfo), error)
}
