package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
)

type account struct {
	name     string
	email    string
	password string
}

type requestBodyCreateAccount struct {
	Name     string
	Email    string
	Password string
}

type requestBodyLoginAccount struct {
	Name     string
	Password string
}

type AccountController struct {
	accountRepo    accountRepository
	tokenGenerator jwt_util.JwtTokenGenerator
}

type AccountInfoDuplicate = int

const (
	UNDEFINED AccountInfoDuplicate = iota
	DUPLICATE_EMAIL
	DUPLICATE_NAME
	NO_DUPLICATES
)

type accountRepository interface {
	createAccount(acc account) error
	checkDuplicate(acc account) (AccountInfoDuplicate, error)
	findAccount(name string) (*account, error)
}

type inMemoryAccountRepository struct {
	accounts []account
}

func (repo *inMemoryAccountRepository) createAccount(acc account) error {
	repo.accounts = append(repo.accounts, acc)

	fmt.Println(repo.accounts)
	return nil
}

func (repo *inMemoryAccountRepository) findAccount(name string) (*account, error) {

	for _, acc := range repo.accounts {
		if acc.name == name {

			// return a pointer to a deep copy of acc, to avoid memory aliasing
			// and giving the caller write access to the repo memory
			return &account{
				name:     strings.Clone(acc.name),
				password: strings.Clone(acc.password),
				email:    strings.Clone(acc.email),
			}, nil
		}
	}
	return nil, nil
}

func (repo *inMemoryAccountRepository) checkDuplicate(acc account) (AccountInfoDuplicate, error) {

	for _, a := range repo.accounts {
		if a.name == acc.name {
			return DUPLICATE_NAME, nil
		}
		if a.email == acc.email {
			return DUPLICATE_EMAIL, nil
		}
	}

	return NO_DUPLICATES, nil

}

func main() {
	var repo inMemoryAccountRepository = inMemoryAccountRepository{
		accounts: make([]account, 0),
	}
	repo.accounts = append(repo.accounts, account{name: "Nele", email: "nele@nele.de", password: "xyz123"})
	repo.accounts = append(repo.accounts, account{name: "Alex", email: "alex@nele.de", password: "abc123"})
	repo.accounts = append(repo.accounts, account{name: "Fabi", email: "fabi@nele.de", password: "def123"})

	tokenGeneratorConfig := jwt_util.JwtConfig{
		SignKey: "jwt_private_key.key",
	}
	tokenGenerator, err := jwt_util.NewJwtTokenGenerator(tokenGeneratorConfig)
	if err != nil {
		panic(err)
	}

	c := AccountController{
		accountRepo:    &repo,
		tokenGenerator: *tokenGenerator,
	}

	fmt.Println("Hello from Auth!")

	r := router.New()

	r.POST("/account", c.handleCreateAccount)
	r.POST("/login", c.handleLogin)

	err = http.ListenAndServe("localhost:8080", r)
	panic(err)
}

func (ctrl *AccountController) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var requestBody requestBodyCreateAccount
	if err := json.Unmarshal(body, &requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if requestBody.Email == "" || requestBody.Name == "" || requestBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newAcc := account{name: requestBody.Name, email: requestBody.Email, password: requestBody.Password}

	duplicate, err := ctrl.accountRepo.checkDuplicate(newAcc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch duplicate {
	case DUPLICATE_NAME:
		w.WriteHeader(http.StatusBadRequest)
	case DUPLICATE_EMAIL:
		w.WriteHeader(http.StatusBadRequest)
	case NO_DUPLICATES:
		err := ctrl.accountRepo.createAccount(newAcc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		panic("unexpected value")
	}

}

func (ctrl *AccountController) handleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestBody requestBodyLoginAccount
	if err := json.Unmarshal(body, &requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestBody.Name
	password := requestBody.Password

	acc, err := ctrl.accountRepo.findAccount(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if acc == nil {
		// username not found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println(acc)

	if acc.password != password {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Falsches Passwort!")
		return
	}

	jwtToken, err := ctrl.tokenGenerator.CreateToken(map[string]interface{}{ //todo: Struct serializen statt map
		"name": acc.name,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(jwtToken)

	cookie := http.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login")

}
