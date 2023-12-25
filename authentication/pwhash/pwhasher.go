package pwhash

type PasswordHasher interface {
	Hash(password string) ([]byte, error)
	Verify(hash []byte, password string) bool
}
