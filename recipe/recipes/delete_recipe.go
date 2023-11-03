package recipes

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"net/http"
)

func (ctrl *Controller) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

	recipeId, err := model.RecipeIdFromString(r.Context().Value("id").(string))

	if err != nil {
		panic(err)
	}

	err = ctrl.repo.DeleteRecipe(recipeId)
	if err != nil {
		fmt.Printf("Failed to delete recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
