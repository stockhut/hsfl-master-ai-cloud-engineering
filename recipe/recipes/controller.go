package recipes

import (
	"encoding/json"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
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
	Name         string                  `json:"name"`
	Ingredients  []ingredientRequestBody `json:"ingredients"`
	Directions   []string                `json:"directions"`
	TimeEstimate int                     `json:"time_estimate"`
	Difficulty   string                  `json:"difficulty"`
	FeedsPeople  int                     `json:"feeds_people"`
}

type recipeResponseModel struct {
	Author       string                   `json:"author"`
	Id           model.RecipeId           `json:"id"`
	Name         string                   `json:"name"`
	Ingredients  []ingredientResponseBody `json:"ingredients"`
	Directions   []string                 `json:"directions"`
	TimeEstimate int                      `json:"time_estimate"`
	Difficulty   string                   `json:"difficulty"`
	FeedsPeople  int                      `json:"feeds_people"`
}

type ingredientRequestBody struct {
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

type ingredientResponseBody struct {
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

func mapx[T any, U any](ts []T, f func(T) U) []U {

	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}

	return us
}

func ingredientRequestToModel(i ingredientRequestBody) model.Ingredient {
	return model.Ingredient{
		Name:   i.Name,
		Unit:   i.Unit,
		Amount: i.Amount,
	}
}

func ingredientModelToResponse(i model.Ingredient) ingredientResponseBody {
	return ingredientResponseBody{
		Name:   i.Name,
		Unit:   i.Unit,
		Amount: i.Amount,
	}
}

func recipeToResponseModel(recipe model.Recipe) recipeResponseModel {
	return recipeResponseModel{
		Id:           recipe.Id,
		Author:       recipe.Author,
		Name:         recipe.Name,
		Ingredients:  mapx(recipe.Ingredients, ingredientModelToResponse),
		Directions:   recipe.Directions,
		TimeEstimate: recipe.TimeEstimate,
		Difficulty:   recipe.Difficulty,
		FeedsPeople:  recipe.FeedsPeople,
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

	recipe := model.Recipe{
		Author:       username.(string),
		Name:         requestBody.Name,
		Ingredients:  mapx(requestBody.Ingredients, ingredientRequestToModel),
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
