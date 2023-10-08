package main

import (
	"fmt"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
)

type account struct {
	name     string
	email    string
	password string
}

var accounts []account = make([]account, 0) //temporary database

func main() {
	accounts = append(accounts, account{name: "Nele", email: "nele@nele.de", password: "xyz123"})
	accounts = append(accounts, account{name: "Alex", email: "alex@nele.de", password: "abc123"})
	accounts = append(accounts, account{name: "Fabi", email: "fabi@nele.de", password: "def123"})

	fmt.Println("Hello from Auth!")

	r := router.New()

	r.GET("/", handleHi)
	r.POST("/createAccount", handleCreateAccount)
	r.GET("/login/:name/:password", handleLogin)

	err := http.ListenAndServe("localhost:8080", r)
	panic(err)
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi!")
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Account Created!")

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
