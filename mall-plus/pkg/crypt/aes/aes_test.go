package aes

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCBC(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("test key")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cbc encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	cipherText, err := CBCEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatalf("aes cbc encrypt err:%v", err)
	}
	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("AES-CBC加密内容: ", cipherTextStr)

	originText, err := CBCDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatalf("aes cbc decrypt err:%v", err)
	}
	t.Log("AES-CBC解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCFB(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cfb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CFBEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("aes cfb encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-CFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := CFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("aes cfb decrypt err: ", err)
	}
	t.Log("AES-CFB解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCTR(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ctr encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CTREncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("aes ctr encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-CTR加密内容: ", cipherTextStr)

	// 解密
	originText, err := CTRDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("aes ctr decrypt err: ", err)
	}
	t.Log("AES-CTR解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFB(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := OFBEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("aes ofb encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-OFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := OFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("aes ofb decrypt err: ", err)
	}
	t.Log("AES-OFB解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBStream(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := OFBEncryptStreamReader(origin, key, iv, false)
	if err != nil {
		t.Fatal("aes ofb stream encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-OFB-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := OFBDecryptStreamWriter(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("aes ofb stream decrypt err: ", err)
	}
	t.Log("AES-OFB-Stream方式解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCBCTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cbc-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	// 加密
	cipherText, err := CBCEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("aes cbc triple encrypt err: ", err)
	}

	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("AES-CBC-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CBCDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("aes cbc triple decrypt err: ", err)
	}
	t.Log("AES-CBC-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCFBTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-cfb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CFBEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("aes cfb triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-CFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("aes cfb triple decrypt err: ", err)
	}
	t.Log("AES-CFB-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCTRTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ctr-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CTREncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("aes ctr triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-CTR-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CTRDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("aes ctr triple decrypt err: ", err)
	}
	t.Log("AES-CTR-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := OFBEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("aes ofb triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-OFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := OFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("aes ofb triple decrypt err: ", err)
	}
	t.Log("AES-OFB-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBStreamTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to aes-ofb-triple-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := OFBEncryptStreamReader(origin, key, iv, true)
	if err != nil {
		t.Fatal("aes ofb stream triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("AES-OFB-Triple-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := OFBDecryptStreamWriter(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("aes ofb stream triple decrypt err: ", err)
	}
	t.Log("AES-OFB-Triple-Stream方式解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}