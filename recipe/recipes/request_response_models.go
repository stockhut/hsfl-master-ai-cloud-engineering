package recipes

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"strconv"
)
import "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"

// The createRecipeRequestBody might be json encoded by the browser in different ways:
//
// - when using forms and the htmx json extension,
// the browser encodes an ingredient into multiple arrays (one for each field) instead a single array of structs
//
// - when there is only one ingredient,
// the Ingredient-* fields may be a single string instead of an array (again, this is done by the browser)
//
// - for other clients, we accept the more convenient array of structs (ingredients) in the json payload
//
// After the JSON is decoded, all ingredients are available in the Ingredients field
type createRecipeRequestBody struct {
	Name              string                  `json:"name"`
	Ingredients       []ingredientRequestBody `json:"ingredients"`
	IngredientNames   any                     `json:"ingredient-name"`
	IngredientUnits   any                     `json:"ingredient-unit"`
	IngredientAmounts any                     `json:"ingredient-amount"`
	Directions        string                  `json:"directions"`
	TimeEstimate      int                     `json:"time_estimate"`
	Difficulty        string                  `json:"difficulty"`
	FeedsPeople       int                     `json:"feeds_people"`
}

func (body *createRecipeRequestBody) UnmarshalJSON(data []byte) error {

	// https://stackoverflow.com/a/50110186

	// to avoid infinite recursion with the json.Unmarshal call below
	type createRecipeRequestBodyTypeCopy createRecipeRequestBody
	var unmarshalTarget createRecipeRequestBodyTypeCopy

	err := json.Unmarshal(data, &unmarshalTarget)
	if err != nil {
		return err
	}
	*body = createRecipeRequestBody(unmarshalTarget)

	names := asStringSlice(unmarshalTarget.IngredientNames)
	amounts := asStringSlice(unmarshalTarget.IngredientAmounts)
	units := asStringSlice(unmarshalTarget.IngredientUnits)

	for i := 0; i < max(len(names), len(amounts), len(units)); i++ {
		amount, err := strconv.Atoi(amounts[i])
		if err != nil {
			return errors.Wrapf(err, "Failed to convert %s to int", amounts[i])
		}
		body.Ingredients = append(body.Ingredients, ingredientRequestBody{
			Name:   names[i],
			Unit:   units[i],
			Amount: amount,
		})
	}

	return nil
}

// asStringSlice converts a single string or a slice of strings into a slice of strings
func asStringSlice(input interface{}) []string {
	a := make([]string, 0)
	switch input.(type) {
	case string:
		a = append(a, input.(string))
	case []interface{}:
		for _, elem := range input.([]interface{}) {
			a = append(a, elem.(string))
		}
	}

	return a
}

type recipeResponseModel struct {
	Author       string                   `json:"author"`
	Id           model.RecipeId           `json:"id"`
	Name         string                   `json:"name"`
	Ingredients  []ingredientResponseBody `json:"ingredients"`
	Directions   string                   `json:"directions"`
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

func recipeRequestToModel(request createRecipeRequestBody, author string) model.Recipe {
	return model.Recipe{
		Author:       author,
		Name:         request.Name,
		Ingredients:  fun.Map(request.Ingredients, ingredientRequestToModel),
		Directions:   request.Directions,
		TimeEstimate: request.TimeEstimate,
		Difficulty:   request.Difficulty,
		FeedsPeople:  request.FeedsPeople,
	}
}
