package recipes

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/html_presenter"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
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
		html_presenter.Present(w, http.StatusOK, ctrl.htmlTemplates, "displayRecipesShort.html", response)
	} else {
		json_presenter.Present(w, http.StatusOK, response)
	}

}
