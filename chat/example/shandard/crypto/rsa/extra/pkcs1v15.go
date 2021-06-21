package extra

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Encrypt(publicKey, originText []byte) ([]byte, error) {
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return nil,err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub, originText)
}

func Decrypt(privateKey, cipherText []byte) ([]byte, error) {
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pri, cipherText)
}

func Sign(privateKey, originText []byte) ([]byte, error) {
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	hashed := sha256.Sum256(originText)

	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hashed[:])
}

func Verify(publicKey, originText, signature []byte) error {
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256(originText)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}