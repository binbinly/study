package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var users sync.Map

	users.Store("pi", 1)
	users.Store("big", 2)
	users.Store("star", 3)

	if u1, ok := users.Load("pi"); ok {
		fmt.Println(u1)
	}

	value, exist := users.LoadOrStore("hello", 66)
	fmt.Println(value, exist)

	users.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	users.Delete("pi")
	fmt.Println("=======delete========")
	users.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}