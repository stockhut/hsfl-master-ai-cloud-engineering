package recipes

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"io"
	"net/http"
)

func (ctrl *Controller) CreateRecipe(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestBody createRecipeRequestBody
	if err := json.Unmarshal(body, &requestBody); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := r.Context().Value(middleware.JwtContextKey).(jwt.MapClaims)

	username, ok := claims["name"]
	if !ok {
		fmt.Println("failed to read name from jwt")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recipe := model.Recipe{
		Author:       username.(string),
		Name:         requestBody.Name,
		Ingredients:  fun.Map(requestBody.Ingredients, ingredientRequestToModel),
		Directions:   requestBody.Directions,
		TimeEstimate: requestBody.TimeEstimate,
		Difficulty:   requestBody.Difficulty,
		FeedsPeople:  requestBody.FeedsPeople,
	}

	newRecipe, err := ctrl.repo.CreateRecipe(recipe)
	if err != nil {
		fmt.Printf("Failed to save recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := recipeToResponseModel(newRecipe)

	responseBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to serialize recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseBytes)
	if err != nil {
		fmt.Printf("failed to write response: %s\n", err)
	}
}
