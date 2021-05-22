package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// raw 使用原生包封装的 http client

// rawClient
type rawClient struct{}

// Get get data by get method
func (r *rawClient) Get(url string, params map[string]string, duration time.Duration) ([]byte, error) {
	client := http.Client{Timeout: duration}
	var target []byte

	resp, err := client.Get(url)
	if err != nil {
		return target, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return target, err
	}

	if err := json.Unmarshal(b, &target); err != nil {
		return target, fmt.Errorf("can't unmarshal to target err: %s, body: %s", err, b)
	}

	return target, nil
}

// Post send data by post method
func (r *rawClient) Post(url string, data []byte, duration time.Duration) ([]byte, error) {
	client := http.Client{Timeout: duration}
	var target []byte
	resp, err := client.Post(url, contentTypeJSON, bytes.NewBuffer(data))
	if err != nil {
		return target, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return target, err
	}

	if err := json.Unmarshal(b, &target); err != nil {
		return target, fmt.Errorf("can't unmarshal to target, err: %s, body: %s", err, b)
	}

	return target, nil
}