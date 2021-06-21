package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
)

type Handle struct{}

func (h *Handle) ServeHTTP(r http.ResponseWriter, request *http.Request) {
	h.Common(r, request)
}

func (h *Handle) Common(r http.ResponseWriter, request *http.Request) {
	hystrix.ConfigureCommand("mycommand", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  20,
		SleepWindow:            5000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  30,
	})
	var msg string

	err := hystrix.Do("mycommand", func() error {
		_, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Printf("请求失败:%v", err)
			msg = "get error"
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Printf("handle  error:%v\n", err)

		return err
	})
	fmt.Printf("err:%v\n", err)
	if err == nil {
		msg = "request success"
	} else {
		msg = fmt.Sprint("get an error, handle it, err:", err)
	}
	r.Write([]byte(msg))
}

func main() {
	http.ListenAndServe(":8090", &Handle{})
}
