package utils

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/teris-io/shortid"
)

var (
	//Hostname 主机名
	Hostname = GetHostname()
)

// GenShortID 生成一个id
func GenShortID() (string, error) {
	return shortid.Generate()
}

// GetBytes interface 转 byte
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const pattern = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	salt := make([]byte, 0, n)
	l := len(pattern)

	for i := 0; i < n; i++ {
		p := r.Intn(l)
		salt = append(salt, pattern[p])
	}

	return string(salt)
}

// RegexpReplace ...
func RegexpReplace(reg, src, temp string) string {
	var result []byte
	pattern := regexp.MustCompile(reg)
	for _, matches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, matches)
	}
	return string(result)
}

// GetHostname 获取主机名
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	return name
}