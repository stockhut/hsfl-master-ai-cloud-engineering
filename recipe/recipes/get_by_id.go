package recipes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
