package recipes

import (
	"encoding/json"
	"fmt"
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
