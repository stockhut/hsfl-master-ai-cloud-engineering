package jwt

type Config interface {
	ReadPrivateKey() (any, error)
}
