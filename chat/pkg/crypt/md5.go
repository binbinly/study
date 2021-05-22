package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 md5加密
func Md5(_, value string) []byte {
	m := md5.New()
	m.Write([]byte(value))

	return m.Sum(nil)
}

//Md5ToString md5加密返回字符串
func Md5ToString(value string) string {
	return hex.EncodeToString(Md5("", value))
}

//Md5ByReader md5加密
func Md5ByReader(r io.Reader) (string, error) {
	m := md5.New()
	_, err := io.Copy(m, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}