package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/pwhash"
	"io"
	"log"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
)

type Controller struct {
	accountRepo    AccountRepository
	tokenGenerator jwt_util.JwtTokenGenerator
	pwHasher       pwhash.PasswordHasher
}

func NewController(accountRepo AccountRepository, tokenGenerator jwt_util.JwtTokenGenerator, pwHaser pwhash.PasswordHasher) *Controller {
	return &Controller{
		accountRepo:    accountRepo,
		tokenGenerator: tokenGenerator,
		pwHasher:       pwHaser,
	}
}

func (ctrl *Controller) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Start HandleCreateAccount")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to read RequestBody: %s\n", err)
		return
	}
	var requestBody requestBodyCreateAccount
	if err := json.Unmarshal(body, &requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to Unmarshal RequestBody: %s\n", err)

		return
	}
	if requestBody.Email == "" || requestBody.Name == "" || requestBody.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Empty RequestBody Email or name or password")
		return
	}

	pwHash, err := ctrl.pwHasher.Hash(requestBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to hash password: %s\n", err)
		return
	}

	newAcc := model.Account{Name: requestBody.Name, Email: requestBody.Email, PasswordHash: pwHash}

	err = ctrl.accountRepo.CheckDuplicate(r.Context(), newAcc)

	if err == nil {
		err := ctrl.accountRepo.CreateAccount(r.Context(), newAcc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		return
	}

	switch {
	case errors.Is(err, ErrDuplicateName):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Username is already taken")
	case errors.Is(err, ErrDuplicateEmail):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Email is already registered")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to check for account duplicate: %s\n", err)
	}

}

func (ctrl *Controller) HandleLogin(w http.ResponseWriter, r *http.Request) {
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

	acc, err := ctrl.accountRepo.FindAccount(r.Context(), username)

	if err != nil {
		log.Printf("Failed to find account %s: %s", username, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if acc == nil {
		// username not found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println(acc)

	if ctrl.pwHasher.Verify(acc.PasswordHash, password) == false {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Falsches Passwort!")
		return
	}

	jwtToken, err := ctrl.tokenGenerator.CreateToken(map[string]interface{}{ //todo: Struct serializen statt map
		"name": acc.Name,
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
