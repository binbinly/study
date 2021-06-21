package extra

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
)

func EncryptOAEP(publicKey, originText, label []byte) ([]byte, error) {
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, originText, label)
}

func DecryptOAEP(privateKey, cipherText, label []byte) ([]byte, error) {
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, pri, cipherText, label)
}

func BuildRSAPublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	return pub, nil
}

func BuildRSAPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pri := priInterface.(*rsa.PrivateKey)

	return pri, nil
}

func BuildRSAPKCS1PublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func BuildRSAPKCS1PrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}