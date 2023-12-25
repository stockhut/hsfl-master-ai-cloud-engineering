package pwhash

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcryptPasswordHasher(t *testing.T) {

	pw := "password"
	t.Run("creates bcrypt hash", func(t *testing.T) {

		bc := BcryptPasswordHasher{}

		hash, err := bc.Hash(pw)

		assert.Nil(t, err)
		// regex based on https://en.wikipedia.org/wiki/Bcrypt#Description
		assert.Regexp(t, "\\$2(a|b|x|y)\\$\\d+\\$.{22}.{31}", string(hash))

		assert.Nil(t, bcrypt.CompareHashAndPassword(hash, []byte(pw)))

	})

	t.Run("verifies it's own hashes", func(t *testing.T) {
		bc := BcryptPasswordHasher{}

		hash, err := bc.Hash(pw)

		assert.Nil(t, err)
		assert.True(t, bc.Verify(hash, pw))
	})

	t.Run("denies verification of bad hashes", func(t *testing.T) {
		bc := BcryptPasswordHasher{}

		hash, err := bc.Hash(pw)

		assert.Nil(t, err)
		assert.False(t, bc.Verify(hash, "not the password"))
	})

}
