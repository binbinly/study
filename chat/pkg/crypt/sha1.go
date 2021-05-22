package crypt

import (
	"crypto/sha1"
	"encoding/hex"
)

//sha1加密
func EncodeSha1(s string) (str string) {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}