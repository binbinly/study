package main

import (
	"fmt"

	"golang.org/x/time/rate"
)

func main()  {

	limit := rate.NewLimiter(50, 100)

	for i := 0; i < 1000; i++ {
		allow := limit.Allow()
		if allow {
			fmt.Printf("i=%d is allow\n", i)
			continue
		}
		fmt.Printf("i=%d is not allow\n", i)
	}
}
