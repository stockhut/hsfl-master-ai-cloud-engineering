package recipes

import (
	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
	"html/template"
	"net/http"
)

func (ctrl *Controller) GetBySelf(w http.ResponseWriter, r *http.Request) {

	// TODO: check if user exists, needs repository for accounts/users

	claims := r.Context().Value(middleware.JwtContextKey).(jwt.MapClaims)

	user, ok := claims["name"].(string)
	if !ok {
		panic("kein name im claim")
	}

	recipes, err := ctrl.repo.GetAllByAuthor(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fun.Map(recipes, recipeToResponseModel)

	if htmx.IsHtmxRequest(r) {
		tmplFile := "templates/displayRecipe.html"
		tmpl, err := template.ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		err = tmpl.Execute(w, response)
		if err != nil {
			panic(err)
		}
	} else {
		json_presenter.JsonPresenter(w, http.StatusOK, response)
	}

}
