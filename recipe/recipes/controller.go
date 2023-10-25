package recipes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/golang-jwt/jwt"
)

type Controller struct {
	repo RecipeRepository
}

func NewController(repo RecipeRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

type createRecipeRequestBody struct {
	Name         string
	Ingredients  []ingredientRequestBody
	Directions   []string
	TimeEstimate int
	Difficulty   string
	FeedsPeople  int
}

type recipeResponseModel struct {
	Author       string
	Id           RecipeId
	Name         string
	Ingredients  []ingredientResponseBody
	Directions   []string
	TimeEstimate int
	Difficulty   string
	FeedsPeople  int
}

type ingredientRequestBody struct {
	Name   string
	Unit   string
	Amount int
}

type ingredientResponseBody struct {
	Name   string
	Unit   string
	Amount int
}

func mapx[T any, U any](ts []T, f func(T) U) []U {

	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}

	return us
}

func ingredientRequestToModel(i ingredientRequestBody) Ingredient {
	return Ingredient{
		name:   i.Name,
		unit:   i.Unit,
		amount: i.Amount,
	}
}

func ingredientModelToResponse(i Ingredient) ingredientResponseBody {
	return ingredientResponseBody{
		Name:   i.name,
		Unit:   i.unit,
		Amount: i.amount,
	}
}

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

	recipe := Recipe{
		author:       username.(string),
		name:         requestBody.Name,
		ingredients:  mapx(requestBody.Ingredients, ingredientRequestToModel),
		directions:   requestBody.Directions,
		timeEstimate: requestBody.TimeEstimate,
		difficulty:   requestBody.Difficulty,
		feedsPeople:  requestBody.FeedsPeople,
	}

	newRecipe, err := ctrl.repo.CreateRecipe(recipe)
	if err != nil {
		fmt.Printf("Failed to save recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := recipeResponseModel{
		Id:           newRecipe.id,
		Author:       newRecipe.author,
		Name:         newRecipe.name,
		Ingredients:  mapx(newRecipe.ingredients, ingredientModelToResponse),
		Directions:   newRecipe.directions,
		TimeEstimate: newRecipe.timeEstimate,
		Difficulty:   newRecipe.difficulty,
		FeedsPeople:  newRecipe.feedsPeople,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to serialize recipe: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBytes)
	if err != nil {
		fmt.Printf("failed to write response: %s\n", err)
	}
}
