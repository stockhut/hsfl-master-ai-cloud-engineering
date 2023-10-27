package recipes

import "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
import "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"

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
		Ingredients:  fun.Map(recipe.Ingredients, ingredientModelToResponse),
		Directions:   recipe.Directions,
		TimeEstimate: recipe.TimeEstimate,
		Difficulty:   recipe.Difficulty,
		FeedsPeople:  recipe.FeedsPeople,
	}
}
