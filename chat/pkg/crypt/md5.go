package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

//Md5 实现了MD5哈希算法
func Md5(value string) []byte {
	// 返回一个新的使用MD5校验的hash.Hash
	m := md5.New()
	// 写入
	m.Write([]byte(value))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	return m.Sum(nil)
}

//Md5ToString 返回md5字符串
func Md5ToString(value string) string {
	return hex.EncodeToString(Md5(value))
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