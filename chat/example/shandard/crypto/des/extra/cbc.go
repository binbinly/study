package extra

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func CBCEncrypt(originText, key, iv []byte, triple bool) ([]byte, error) {
	var block cipher.Block

	var err error

	if triple {
		block, err = des.NewTripleDESCipher(key)
	} else {
		block, err = des.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	originText = PKCS5Padding(originText, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(originText))

	blockMode.CryptBlocks(cipherText, originText)

	return cipherText, nil
}

func CBCDecrypt(cipherText, key, iv []byte, triple bool) ([]byte, error) {
	var block cipher.Block

	var err error

	if triple {
		block, err = des.NewTripleDESCipher(key)
	} else {
		block, err = des.NewCipher(key)
	}

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)

	originText := make([]byte, len(cipherText))

	blockMode.CryptBlocks(originText, cipherText)

	originText = PKCS5UnPadding(originText)

	return originText, nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	less := len(cipherText) % blockSize

	padding := blockSize - less

	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func PKCS5UnPadding(originText []byte) []byte {

	length := len(originText)

	unPadding := int(originText[length-1])

	return originText[:(length - unPadding)]
}