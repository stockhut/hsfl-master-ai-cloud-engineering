package main

import (
	"bytes"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/pwhash"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/environment"
	requestlogger "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request-logger"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
)

type inMemoryAccountRepository struct {
	accounts []model.Account
}

func (repo *inMemoryAccountRepository) CreateAccount(acc model.Account) error {
	repo.accounts = append(repo.accounts, acc)

	fmt.Println(repo.accounts)
	return nil
}

func (repo *inMemoryAccountRepository) FindAccount(name string) (*model.Account, error) {

	for _, acc := range repo.accounts {
		if acc.Name == name {

			// return a pointer to a deep copy of acc, to avoid memory aliasing
			// and giving the caller write access to the repo memory
			return &model.Account{
				Name:         strings.Clone(acc.Name),
				PasswordHash: bytes.Clone(acc.PasswordHash),
				Email:        strings.Clone(acc.Email),
			}, nil
		}
	}
	return nil, nil
}

func (repo *inMemoryAccountRepository) CheckDuplicate(acc model.Account) (repository.AccountInfoDuplicate, error) {

	for _, a := range repo.accounts {
		if a.Name == acc.Name {
			return repository.DUPLICATE_EMAIL, nil
		}
		if a.Email == acc.Email {
			return repository.DUPLICATE_EMAIL, nil
		}
	}

	return repository.NO_DUPLICATES, nil

}

const JwtPrivateKeyEnvKey = "JWT_PRIVATE_KEY"

func main() {

	jwtPrivateKeyFile := environment.GetRequiredEnvVar(JwtPrivateKeyEnvKey)

	var repo inMemoryAccountRepository = inMemoryAccountRepository{
		accounts: make([]model.Account, 0),
	}
	repo.accounts = append(repo.accounts, model.Account{Name: "Nele", Email: "nele@nele.de", PasswordHash: []byte("$2y$10$jizNLzGmDsmn05gTH4qwsO8FW2T249SKrpKKeUvUsVfyJUQXbZZra")})
	repo.accounts = append(repo.accounts, model.Account{Name: "Alex", Email: "alex@nele.de", PasswordHash: []byte("$2y$10$jizNLzGmDsmn05gTH4qwsO8FW2T249SKrpKKeUvUsVfyJUQXbZZra")})
	repo.accounts = append(repo.accounts, model.Account{Name: "Fabi", Email: "fabi@nele.de", PasswordHash: []byte("$2y$10$jizNLzGmDsmn05gTH4qwsO8FW2T249SKrpKKeUvUsVfyJUQXbZZra")})

	tokenGeneratorConfig := jwt_util.JwtConfig{
		SignKey: jwtPrivateKeyFile,
	}
	tokenGenerator, err := jwt_util.NewJwtTokenGenerator(tokenGeneratorConfig)
	if err != nil {
		panic(err)
	}

	bcryptPwHasher := pwhash.BcryptPasswordHasher{}
	c := accounts.NewController(&repo, *tokenGenerator, &bcryptPwHasher)

	fmt.Println("Hello from Auth!")

	logFlags := log.Ltime | log.Lmsgprefix | log.Lmicroseconds
	logger := log.New(os.Stdout, "", logFlags)
	logMw := requestlogger.New(logger)

	r := router.New(logMw, c)

	err = http.ListenAndServe(":8080", r)
	panic(err)
}
