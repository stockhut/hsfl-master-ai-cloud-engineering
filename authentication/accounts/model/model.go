package model

type Account struct {
	Name         string
	Email        string
	PasswordHash []byte
}
