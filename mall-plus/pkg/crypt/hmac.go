package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

// Hmac 实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）
func Hmac(key, content string) []byte {
	// 返回一个采用hash.Hash作为底层hash接口、key作为密钥的HMAC算法的hash接口
	h := hmac.New(sha256.New, []byte(key))
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	expectedMAC := h.Sum(nil)
	return expectedMAC
}
