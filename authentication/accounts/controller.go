package accounts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/model"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
)

type Controller struct {
	accountRepo    AccountRepository
	tokenGenerator jwt_util.JwtTokenGenerator
}

func NewController(accountRepo AccountRepository, tokenGenerator jwt_util.JwtTokenGenerator) *Controller {
	return &Controller{
		accountRepo:    accountRepo,
		tokenGenerator: tokenGenerator,
	}
}

func (ctrl *Controller) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
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
	newAcc := model.Account{Name: requestBody.Name, Email: requestBody.Email, Password: requestBody.Password}

	duplicate, err := ctrl.accountRepo.CheckDuplicate(newAcc)
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
		err := ctrl.accountRepo.CreateAccount(newAcc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		panic("unexpected value")
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

	acc, err := ctrl.accountRepo.FindAccount(username)

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

	if acc.Password != password {
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
