package main

import (
	"chat/example/shandard/crypto/aes/extra"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func main()  {
	CBC()

	GCM()

	CFB()

	CTR()

	OFB()

	OFBStream()
}

func CBC() {

	// 声明一个16字节的key
	var key = []byte("example key 1234")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cbc encode test text")
	// 声明一个16字节的iv
	var iv = []byte("example iv tests")

	// 加密
	cipherText, err := extra.CBCEncrypt(origin, key, iv)
	if err != nil {
		log.Fatal(err)
	}

	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("AES-CBC加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CBCDecrypt(cipherText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-CBC解密内容: ", string(originText))
}

func GCM() {

	// 声明一个16字节的key
	var key = []byte("0123456789ABCDEF")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-gcm encode test text")

	// 初始化一个长度为12字节的空的[]byte，不要使用超过2^32个随机非字符，因为存在重复的风险
	nonce := make([]byte, 12)
	// 使用rand随机生成数据
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	// 加密
	cipherText, err := extra.GCMEncrypt(origin, key, nonce)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("AES-GCM加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.GCMDecrypt(cipherText, key, nonce)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-GCM解密内容: ", string(originText))
}

func CFB() {

	// 声明一个16字节的key
	var key = []byte("0123456789ABCDEF")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cfb encode test text")
	// 声明一个16字节的iv
	var iv = []byte("example iv tests")

	// 加密
	cipherText, err := extra.CFBEncrypt(origin, key, iv)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("AES-CFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CFBDecrypt(cipherText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-CFB解密内容: ", string(originText))
}

func CTR() {

	// 声明一个16字节的key
	var key = []byte("0123456789ABCDEF")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ctr encode test text")
	// 声明一个16字节的iv
	var iv = []byte("example iv tests")

	// 加密
	cipherText, err := extra.CTREncrypt(origin, key, iv)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("AES-CTR加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.CTRDecrypt(cipherText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-CTR解密内容: ", string(originText))
}

func OFB() {

	// 声明一个16字节的key
	var key = []byte("0123456789ABCDEF")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb encode test text")
	// 声明一个16字节的iv
	var iv = []byte("example iv tests")

	// 加密
	cipherText, err := extra.OFBEncrypt(origin, key, iv)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("AES-OFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := extra.OFBDecrypt(cipherText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-OFB解密内容: ", string(originText))
}

func OFBStream() {

	// 声明一个16字节的key
	var key = []byte("0123456789ABCDEF")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb-stream encode test text")
	// 声明一个16字节的iv
	var iv = []byte("example iv tests")

	// StreamReader方式加密
	cipherText, err := extra.OFBEncryptStreamReader(origin, key, iv)
	if err != nil {
		log.Fatal(err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	fmt.Println("AES-OFB-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := extra.OFBDecryptStreamWriter(cipherText, key, iv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AES-OFB-Stream方式解密内容: ", string(originText))
}

