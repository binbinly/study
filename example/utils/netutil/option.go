package netutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Option is for adding http.Option
type Option func(req *http.Request)

func SetMethod(method string) Option {
	return func(req *http.Request) {
		req.Method = method
	}
}

// SetHeader set retry context config
func SetHeader(header http.Header) Option {
	return func(req *http.Request) {
		for k, vv := range header {
			for _, vvv := range vv {
				req.Header.Add(k, vvv)
			}
		}
	}
}

func SetHeaderMap(header map[string]string) Option {
	return func(req *http.Request) {
		for k := range header {
			req.Header.Add(k, header[k])
		}
	}
}

func SetUrl(reqUrl string) Option {
	return func(req *http.Request) {
		u, err := url.Parse(reqUrl)
		if err != nil {
			panic(fmt.Sprintf("Option url parse err %v", err))
		}
		req.URL = u
	}
}

func SetParam(reqUrl string, values url.Values) Option {
	return func(req *http.Request) {
		if values != nil {
			if !strings.Contains(reqUrl, "?") {
				reqUrl = reqUrl + "?" + values.Encode()
			} else {
				reqUrl = reqUrl + "&" + values.Encode()
			}
		}
		u, err := url.Parse(reqUrl)
		if err != nil {
			panic(fmt.Sprintf("Option url parse err %v", err))
		}
		req.URL = u
	}
}

func SetBody(data []byte) Option {
	return func(req *http.Request) {
		req.Method = http.MethodPost
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
		req.ContentLength = int64(len(data))
	}
}

func SetForm(values url.Values) Option {
	return func(req *http.Request) {
		req.Method = http.MethodPost
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		data := values.Encode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
		req.ContentLength = int64(len(data))
	}
}

func BuildValues(params map[string]interface{}) url.Values {
	values := url.Values{}
	for k := range params {
		values.Set(k, fmt.Sprintf("%v", params[k]))
	}
	return values
}