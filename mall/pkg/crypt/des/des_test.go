package des

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
	var origin = []byte("need to des-cbc encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	cipherText, err := CBCEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatalf("des cbc encrypt err:%v", err)
	}
	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("DES-CBC加密内容: ", cipherTextStr)

	originText, err := CBCDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatalf("des cbc decrypt err:%v", err)
	}
	t.Log("DES-CBC解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCFB(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cfb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CFBEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("des cfb encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-CFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := CFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("des cfb decrypt err: ", err)
	}
	t.Log("DES-CFB解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCTR(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ctr encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CTREncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("des ctr encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-CTR加密内容: ", cipherTextStr)

	// 解密
	originText, err := CTRDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("des ctr decrypt err: ", err)
	}
	t.Log("DES-CTR解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFB(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := OFBEncrypt(origin, key, iv, false)
	if err != nil {
		t.Fatal("des ofb encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-OFB加密内容: ", cipherTextStr)

	// 解密
	originText, err := OFBDecrypt(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("des ofb decrypt err: ", err)
	}
	t.Log("DES-OFB解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBStream(t *testing.T) {
	// 声明一个8字节的key
	var key = []byte("12345678")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := OFBEncryptStreamReader(origin, key, iv, false)
	if err != nil {
		t.Fatal("des ofb stream encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-OFB-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := OFBDecryptStreamWriter(cipherText, key, iv, false)
	if err != nil {
		t.Fatal("des ofb stream decrypt err: ", err)
	}
	t.Log("DES-OFB-Stream方式解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCBCTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cbc-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("test ivs")

	// 加密
	cipherText, err := CBCEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("des cbc triple encrypt err: ", err)
	}

	// byte转base64字符串
	cipherTextStr := base64.StdEncoding.EncodeToString(cipherText)
	t.Log("DES-CBC-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CBCDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("des cbc triple decrypt err: ", err)
	}
	t.Log("DES-CBC-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCFBTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-cfb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CFBEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("des cfb triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-CFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("des cfb triple decrypt err: ", err)
	}
	t.Log("DES-CFB-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestCTRTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ctr-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := CTREncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("des ctr triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-CTR-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := CTRDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("des ctr triple decrypt err: ", err)
	}
	t.Log("DES-CTR-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-triple encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// 加密
	cipherText, err := OFBEncrypt(origin, key, iv, true)
	if err != nil {
		t.Fatal("des ofb triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-OFB-Triple加密内容: ", cipherTextStr)

	// 解密
	originText, err := OFBDecrypt(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("des ofb triple decrypt err: ", err)
	}
	t.Log("DES-OFB-Triple解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}

func TestOFBStreamTriple(t *testing.T) {

	// 声明一个24字节的key
	var key = []byte("it is 24 bytes test key!")
	// 声明一个随意长度的 需加密内容
	var origin = []byte("need to des-ofb-triple-stream encode test text")
	// 声明一个8字节的iv
	var iv = []byte("iv tests")

	// StreamReader方式加密
	cipherText, err := OFBEncryptStreamReader(origin, key, iv, true)
	if err != nil {
		t.Fatal("des ofb stream triple encrypt err: ", err)
	}

	// byte转十六进制字符串
	cipherTextStr := hex.EncodeToString(cipherText)
	t.Log("DES-OFB-Triple-Stream方式加密内容: ", cipherTextStr)

	// StreamWriter方式解密
	originText, err := OFBDecryptStreamWriter(cipherText, key, iv, true)
	if err != nil {
		t.Fatal("des ofb stream triple decrypt err: ", err)
	}
	t.Log("DES-OFB-Triple-Stream方式解密内容: ", string(originText))
	assert.Equal(t, originText, origin)
}