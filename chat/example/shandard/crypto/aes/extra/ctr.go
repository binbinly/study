package extra

import (
	"crypto/aes"
	"crypto/cipher"
)

func CTREncrypt(originText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize + len(originText))

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func CTRDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}