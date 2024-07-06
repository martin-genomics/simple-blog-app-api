package tools

import (
	"crypto/hmac"
	"crypto/sha256"

	"github.com/golang-jwt/jwt/v5"
)

type CustomSigningMethod struct{}

func (m *CustomSigningMethod) Alg() string {
	return "Custom"
}

func (m *CustomSigningMethod) Verify(signingString, signature string, key interface{}) error {
	expectedMAC := computeHMAC(signingString, key.([]byte))
	if !hmac.Equal([]byte(signature), expectedMAC) {
		return jwt.ErrSignatureInvalid
	}
	return nil
}

func (m *CustomSigningMethod) Sign(signingString string, key interface{}) ([]byte, error) {
	sig := computeHMAC(signingString, key.([]byte))
	return sig, nil
}

func computeHMAC(signingString string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signingString))
	return h.Sum(nil)
}
