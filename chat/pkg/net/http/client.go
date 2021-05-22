package http

import (
	"time"
)

// see: https://github.com/iiinsomnia/gochat/blob/master/utils/http.go

const (
	contentTypeJSON = "application/json"
)

// Client 定义 http client 接口
type Client interface {
	Get(url string, params map[string]string, duration time.Duration) ([]byte, error)
	Post(url string, data []byte, duration time.Duration) ([]byte, error)
}

// 实例化 http client
func NewRawClient() Client {
	return &rawClient{}
}

func NewRestyClient() Client {
	return &restyClient{}
}