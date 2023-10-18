package jwt_util

type Config interface {
	ReadPrivateKey() (any, error)
}
