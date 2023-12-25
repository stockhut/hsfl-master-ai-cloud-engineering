package accounts

import (
	"context"
	"errors"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
)

var ErrDuplicateName = errors.New("duplicate account name")
var ErrDuplicateEmail = errors.New("duplicate email")

type AccountRepository interface {
	CreateAccount(ctx context.Context, acc model.Account) error
	CheckDuplicate(ctx context.Context, acc model.Account) error
	FindAccount(ctx context.Context, name string) (*model.Account, error)
}
