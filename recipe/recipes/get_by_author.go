package recipes

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"
	"net/http"
)

func (ctrl *Controller) GetByAuthor(w http.ResponseWriter, r *http.Request) {

	// TODO: check if user exists, needs repository for accounts/users

	author := r.Context().Value("author").(string)

	recipes, err := ctrl.repo.GetAllByAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fun.Map(recipes, recipeToResponseModel)

	json_presenter.JsonPresenter(w, http.StatusOK, response)
}
