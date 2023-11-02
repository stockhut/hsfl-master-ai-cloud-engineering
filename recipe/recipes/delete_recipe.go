package recipes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
)

func (ctrl *Controller) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

	recipeId := r.Context().Value("id").(string)

	id, err := strconv.Atoi(recipeId)

	if err != nil {
		panic(err)
	}

	err = ctrl.repo.DeleteRecipe(model.RecipeId(id))
	if err != nil {
		fmt.Printf("Failed to delete recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
