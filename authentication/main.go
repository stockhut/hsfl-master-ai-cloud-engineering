package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"golang.org/x/exp/slices"
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

var accounts []account = make([]account, 0) //temporary database

func main() {
	accounts = append(accounts, account{name: "Nele", email: "nele@nele.de", password: "xyz123"})
	accounts = append(accounts, account{name: "Alex", email: "alex@nele.de", password: "abc123"})
	accounts = append(accounts, account{name: "Fabi", email: "fabi@nele.de", password: "def123"})

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
	index := slices.IndexFunc(accounts, func(acc account) bool {
		return (acc.email == requestBody.Email || acc.name == requestBody.Name)
	})

	if index < 0 {
		accounts = append(accounts, account{name: requestBody.Name, email: requestBody.Email, password: requestBody.Password})
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	for _, acc := range accounts {
		if r.Context().Value("name") == acc.name {
			if r.Context().Value("password") == acc.password {
				fmt.Fprintln(w, "Login")
			} else {
				fmt.Fprintln(w, "Falsches Passwort!")
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
