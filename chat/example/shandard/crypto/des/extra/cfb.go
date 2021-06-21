package extra

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"github.com/pkg/errors"
)

func CFBEncrypt(originText, key, iv []byte, triple bool) ([]byte, error) {
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

	cipherText := make([]byte, aes.BlockSize + len(originText))

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func CFBDecrypt(cipherText, key, iv []byte, triple bool) ([]byte, error) {
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

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}