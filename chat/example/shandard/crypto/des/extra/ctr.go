package extra

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

func CTREncrypt(originText, key, iv []byte, triple bool) ([]byte, error) {
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

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func CTRDecrypt(cipherText, key, iv []byte, triple bool) ([]byte, error) {
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
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}