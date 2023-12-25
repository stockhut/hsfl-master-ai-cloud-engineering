package pwhash

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHasher struct {
}

func (h *BcryptPasswordHasher) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (h *BcryptPasswordHasher) Verify(hash []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}
