package crypt

import (
	"crypto/hmac"
	"crypto/sha1"
)

// Hmac hmac
func Hmac(secretKey, body string) []byte {
	m := hmac.New(sha1.New, []byte(secretKey))
	_, _ = m.Write([]byte(body))
	return m.Sum(nil)
}
