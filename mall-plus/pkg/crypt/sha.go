package crypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

//Sha sha哈希算法
func Sha(hash func() hash.Hash, content string) string {
	h := hash()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}

//Sha1 实现了SHA1哈希算法
func Sha1(content string) string {
	// 返回一个新的使用SHA1校验的hash.Hash
	h := sha1.New()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}

//Sha224 实现了SHA224哈希算法
func Sha224(content string) string {
	// 返回一个新的使用SHA224校验的hash.Hash
	h := sha256.New224()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}

//Sha256 实现了SHA256哈希算法
func Sha256(content string) string {
	// 返回一个新的使用SHA256校验的hash.Hash
	h := sha256.New()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}

//Sha384 实现了SHA384哈希算法
func Sha384(content string) string {
	// 返回一个新的使用SHA384校验的hash.Hash
	h := sha512.New384()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}

//Sha512 实现了SHA512哈希算法
func Sha512(content string) string {
	// 返回一个新的使用SHA512校验的hash.Hash
	h := sha512.New()
	// 写入
	h.Write([]byte(content))
	// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	m := h.Sum(nil)
	return hex.EncodeToString(m)
}