package extra

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
)

func OFBEncrypt(originText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize + len(originText))

	stream := cipher.NewOFB(block, iv)

	stream.XORKeyStream(cipherText[aes.BlockSize:], originText)

	return cipherText, nil
}

func OFBDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewOFB(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func OFBEncryptStreamReader(originText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)

	reader := &cipher.StreamReader{
		S: stream,
		R: bytes.NewReader(originText),
	}

	return ioutil.ReadAll(reader)
}

func OFBDecryptStreamWriter(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)

	var originText bytes.Buffer

	writer := &cipher.StreamWriter{
		S: stream,
		W: &originText,
	}

	if _, err = io.Copy(writer, bytes.NewReader(cipherText)); err != nil {
		return nil, err
	}
	return originText.Bytes(), nil
}