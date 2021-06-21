package extra

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

func SignPass(privateKey, originText []byte, opts *rsa.PSSOptions) ([]byte, error) {
	pri, err := BuildRSAPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	hash := crypto.MD5
	h := hash.New()
	h.Write(originText)
	hashed := h.Sum(nil)

	return rsa.SignPSS(rand.Reader, pri, hash, hashed, opts)
}

func VerifyPass(publicKey, originText, signature []byte, opts *rsa.PSSOptions) error {
	pub, err := BuildRSAPublicKey(publicKey)
	if err != nil {
		return err
	}

	hash := crypto.MD5

	h := hash.New()

	h.Write(originText)

	hashed := h.Sum(nil)

	return rsa.VerifyPSS(pub, hash, hashed, signature, opts)
}