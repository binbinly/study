package extra

import (
	"crypto/aes"
	"crypto/cipher"
)

func GCMEncrypt(originText, key, nonce []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	g, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := g.Seal(nil, nonce, originText, nil)

	return cipherText, nil
}

func GCMDecrypt(cipherText, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGom, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesGom.Open(nil, nonce, cipherText, nil)
}
