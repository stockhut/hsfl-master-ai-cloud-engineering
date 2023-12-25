package accounts

import (
	"context"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
)

type AccountInfoDuplicate = int

const (
	UNDEFINED AccountInfoDuplicate = iota
	DUPLICATE_EMAIL
	DUPLICATE_NAME
	NO_DUPLICATES
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, acc model.Account) error
	CheckDuplicate(ctx context.Context, acc model.Account) (AccountInfoDuplicate, error)
	FindAccount(ctx context.Context, name string) (*model.Account, error)
}
