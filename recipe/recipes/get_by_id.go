package recipes

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"net/http"
)

func (ctrl *Controller) GetById(w http.ResponseWriter, r *http.Request) {

	recipeId := r.Context().Value("id").(string)

	recipe, err := ctrl.repo.GetById(model.RecipeId(recipeId))
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
