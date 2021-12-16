package main

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main()  {
	err := status.New(codes.ResourceExhausted, "Request limit exceeded.").Err()

	st := status.Convert(err)
	fmt.Println("st", st.Proto())
}