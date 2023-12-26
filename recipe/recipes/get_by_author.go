package recipes

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/html_presenter"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

func (ctrl *Controller) GetByAuthor(w http.ResponseWriter, r *http.Request) {

	// TODO: check if user exists, needs repository for accounts/users

	author := r.Context().Value("author").(string)

	recipes, err := ctrl.repo.GetAllByAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// template.Must(template.ParseGlob("templates/*.gohtml")) <--alex sagt
	response := fun.Map(recipes, recipeToResponseModel)

	if htmx.IsHtmxRequest(r) {
		html_presenter.Present(w, http.StatusOK, ctrl.htmlTemplates, "displayRecipesShort.html", response)
	} else {
		json_presenter.JsonPresenter(w, http.StatusOK, response)
	}
}
