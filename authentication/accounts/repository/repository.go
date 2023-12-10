package repository

import (
	"errors"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
)

type AccountInfoDuplicate = int

const (
	UNDEFINED AccountInfoDuplicate = iota
	DUPLICATE_EMAIL
	DUPLICATE_NAME
	NO_DUPLICATES
)

var ErrAccountNotFound = errors.New("no such account")

type AccountRepository interface {
	CreateAccount(acc model.Account) error
	CheckDuplicate(acc model.Account) (AccountInfoDuplicate, error)
	FindAccount(name string) (*model.Account, error)
}
