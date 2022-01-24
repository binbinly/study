package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger(WithCallerSkipCount(1))
	l.Fields(map[string]interface{}{"key3": "val4"}).Log(InfoLevel, "test_msg")
	Info("test info")
	Warn("test info")
	Error("test info")
}