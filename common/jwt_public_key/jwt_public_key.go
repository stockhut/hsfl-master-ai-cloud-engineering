package jwt_public_key

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

type JwtPublicKey any

func FromFile(path string) (JwtPublicKey, error) {

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fromBytes(bytes)
}

func fromBytes(bytes []byte) (JwtPublicKey, error) {
	block, _ := pem.Decode(bytes)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
