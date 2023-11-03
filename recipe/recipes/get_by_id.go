package recipes

import (
	"net/http"
	"strconv"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
)

func (ctrl *Controller) GetById(w http.ResponseWriter, r *http.Request) {

	recipeId := r.Context().Value("id").(string)

	id, err := strconv.Atoi(recipeId)

	if err != nil {
		panic(err)
	}

	recipe, err := ctrl.repo.GetById(model.RecipeId(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if recipe == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := recipeToResponseModel(*recipe)

	json_presenter.JsonPresenter(w, http.StatusOK, response)
}
