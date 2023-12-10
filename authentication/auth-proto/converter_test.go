package auth_proto

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountResponseFromModel(t *testing.T) {

	acc := model.Account{
		Name:     "bob",
		Email:    "bob@example.org",
		Password: "",
	}

	resp := AccountResponseFromModel(&acc)

	assert.Equal(t, acc.Name, resp.Name)
	assert.Equal(t, acc.Email, resp.Email)
}
