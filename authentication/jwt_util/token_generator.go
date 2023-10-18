package jwt_util

type TokenGenerator interface {
	CreateToken(claims map[string]interface{}) (string, error)
}
