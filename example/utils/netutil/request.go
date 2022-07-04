package netutil

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	timeout = 10 * time.Second
)

type Request struct {
	Debug   bool
	httpReq *http.Request
	Client  *http.Client
}

type Requests func(req *Request)

func WithDebug(debug bool) Requests {
	return func(req *Request) {
		req.Debug = debug
	}
}

func WithClient(client *http.Client) Requests {
	return func(req *Request) {
		req.Client = client
	}
}

func NewRequest(opts ...Requests) *Request {

	req := new(Request)

	req.Client = http.DefaultClient
	req.Client.Timeout = timeout

	for _, opt := range opts {
		opt(req)
	}

	return req
}

func (r *Request) Do(opts ...Option) (*http.Response, error) {
	r.httpReq = &http.Request{
		Method:     http.MethodGet,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	for _, opt := range opts {
		opt(r.httpReq)
	}

	r.debug()
	resp, e := r.Client.Do(r.httpReq)
	return resp, e
}

func (r *Request) debug() {

	if r.Debug == false {
		return
	}

	message, err := httputil.DumpRequestOut(r.httpReq, false)
	if err != nil {
		return
	}
	log.Println(string(message))
}
