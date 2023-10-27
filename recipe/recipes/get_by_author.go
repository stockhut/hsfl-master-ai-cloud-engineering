package recipes

import (
	"encoding/json"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
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

	responseBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to serialize recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBytes)
	if err != nil {
		fmt.Printf("failed to write response: %s\n", err)
	}
}
