package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

var repo inMemoryAccountRepository = inMemoryAccountRepository{
	accounts: make([]account, 0),
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
	findAccount(name string)
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

	repo.accounts = append(repo.accounts, account{name: "Nele", email: "nele@nele.de", password: "xyz123"})
	repo.accounts = append(repo.accounts, account{name: "Alex", email: "alex@nele.de", password: "abc123"})
	repo.accounts = append(repo.accounts, account{name: "Fabi", email: "fabi@nele.de", password: "def123"})

	fmt.Println("Hello from Auth!")

	r := router.New()

	r.POST("/account", handleCreateAccount)
	r.GET("/login/:name/:password", handleLogin)

	err := http.ListenAndServe("localhost:8080", r)
	panic(err)
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
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

	duplicate, err := repo.checkDuplicate(newAcc)
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
		err := repo.createAccount(newAcc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		panic("unexpected value")
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("name").(string)
	password := r.Context().Value("password")

	fmt.Printf("username: %s, password: %s\n", username, password)
	acc, err := repo.findAccount(username)

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

	if acc.password == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Login")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Falsches Passwort!")
	}

}
