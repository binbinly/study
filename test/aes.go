package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
)

// 9OCkv4AMKkzCnJgV40mzzpiChtCEJpU/lz5pcLvapJukMy3W/4xUVQzV4o15wU5Y4pDddQ0s9Uz5FRkrk6AdPUXzi+eTOb4satRfMCeGwAHIyYamDP7+mJyhMNrSx5rl
// W652fJc832q3euZsAUEYtI75Lm7OwiCJ4ggCJU/8OW5kcLoCLsnyXr20ZJm9yKzchTH2teF/kjh+O1vXbkN97T0TJftUbhyRLEvDLr2+x9WtdNo/vAG1C7Ss0aVUjoUN
// 8KDdfxmZJuG/yqW9KiC/MdNNEIIFbdB/Pax7KDHfBRVdFExoR3d0evahQe4bNxsslzernFS5gKvdAsklveRXNC+EFtMKg2sQADcej1+ggtArTEZm49TzuthxpIRkTmwN
// hj1LMczLOdLUkuN89h7dYoeQyxXf/CWIDGMulldnFUzgptwie5HOwvUNelSyda3IJsX6wblAX3OuypMcVAMLPNN8zDkKnVy3Ge7ABwiTAnm0PekuxBN6m/YATDj+l3O1

//vGo0HjQAh8fvMs8IBWa63Cd5ZwRjpJv3HIr4cUf5eQ/rcyN6p28aHonnwuAyX5iM6c3j/crWZ1ZWHfLV6NsZEJmuPlOxj0cFi4jGnMIW61cHiLMOmuxuzHUDlaSMX+CrYcNtKtjHDh73D83N/w0GU68LA+iAFC5jASaiv5VImOhGsUBxUyQhvdKx5z4YJqaSM4TcOgN4CoFfS4kjNFXiprOjRKIKNDRlmCPgYYDT4tev2WQorXZfwMcn+a08CMc/
func main() {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	log.Println("------------------ CBC模式 --------------------")
	encrypted := AesEncryptCBC(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	log.Println("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(encrypted, key)
	log.Println("解密结果：", string(decrypted))
}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
