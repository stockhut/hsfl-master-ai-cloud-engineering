package jwt_util

import (
    "crypto/x509"
	"encoding/pem"
	"os"
)

func ReadPrivateKey(file string) (any, error){
	bytes, err := os.ReadFile("jwt_public_key.key")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(bytes)
	return x509.ParsePKIXPublicKey(block.Bytes)
}