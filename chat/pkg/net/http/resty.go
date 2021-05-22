package http

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// docs: https://github.com/go-resty/resty

type restyClient struct{}

// Get request url by get method
func (r *restyClient) Get(url string, params map[string]string, duration time.Duration) ([]byte, error) {
	client := resty.New()

	if duration != 0 {
		client.SetTimeout(duration)
	}

	if len(params) > 0 {
		client.SetQueryParams(params)
	}

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Content-Type": contentTypeJSON,
		}).
		Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

// Post request url by post method
func (r *restyClient) Post(url string, data []byte, duration time.Duration) ([]byte, error) {
	client := resty.New()

	if duration != 0 {
		client.SetTimeout(duration)
	}

	cr := client.R().
		SetBody(string(data)).
		SetHeaders(map[string]string{
			"Content-Type": contentTypeJSON,
		})

	resp, err := cr.Post(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}