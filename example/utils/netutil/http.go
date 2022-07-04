package netutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

//HttpGet send get http request
func HttpGet(url string, opts ...Option) (*http.Response, error) {
	req := NewRequest()
	opts = append(opts, SetUrl(url))
	return req.Do(opts...)
}

//HttpPost send post http request
func HttpPost(url string, opts ...Option) (*http.Response, error) {
	req := NewRequest()
	opts = append(opts, SetUrl(url))
	opts = append(opts, SetMethod(http.MethodPost))
	return req.Do(opts...)
}

//HttpPut send put http request
func HttpPut(url string, opts ...Option) (*http.Response, error) {
	req := NewRequest()
	opts = append(opts, SetUrl(url))
	opts = append(opts, SetMethod(http.MethodPut))
	return req.Do(opts...)
}

//HttpDelete send delete http request
func HttpDelete(url string, opts ...Option) (*http.Response, error) {
	req := NewRequest()
	opts = append(opts, SetUrl(url))
	opts = append(opts, SetMethod(http.MethodDelete))
	return req.Do(opts...)
}

// HttpPatch send patch http request
func HttpPatch(url string, opts ...Option) (*http.Response, error) {
	req := NewRequest()
	opts = append(opts, SetUrl(url))
	opts = append(opts, SetMethod(http.MethodPatch))
	return req.Do(opts...)
}

// ParseHttpResponse decode http response to specified interface
func ParseHttpResponse(resp *http.Response, obj interface{}) error {
	if resp == nil {
		return errors.New("InvalidResp")
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}

// ConvertMapToQueryString convert map to sorted url query string
func ConvertMapToQueryString(param map[string]interface{}) string {
	if param == nil {
		return ""
	}
	var keys []string
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var build strings.Builder
	for i, v := range keys {
		build.WriteString(v)
		build.WriteString("=")
		build.WriteString(fmt.Sprintf("%v", param[v]))
		if i != len(keys)-1 {
			build.WriteString("&")
		}
	}
	return build.String()
}