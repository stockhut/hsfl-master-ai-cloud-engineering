package auth_proto

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
)

// AccountResponseFromModel converts a model.Account to a GetAccountResponse
func AccountResponseFromModel(acc *model.Account) *GetAccountResponse {
	return &GetAccountResponse{
		Name:  acc.Name,
		Email: acc.Email,
	}
}
