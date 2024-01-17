package recipes

import (
	"errors"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/html_presenter"
	"net/http"
	"strconv"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
)

func (ctrl *Controller) GetById(w http.ResponseWriter, r *http.Request) {

	recipeId := r.Context().Value("id").(string)

	id, err := strconv.Atoi(recipeId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipe, err := ctrl.repo.GetById(model.RecipeId(id))

	if err != nil {
		if errors.Is(err, ErrNoSuchID) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if recipe == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := recipeToResponseModel(*recipe)

	if htmx.IsHtmxRequest(r) {
		html_presenter.Present(w, http.StatusOK, ctrl.htmlTemplates, "displayRecipe.html", response)
	} else {
		json_presenter.Present(w, http.StatusOK, response)
	}
}
