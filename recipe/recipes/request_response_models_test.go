package recipes

import (
	"testing"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
)

func Test_ingredientModelToResponse(t *testing.T) {

	t.Run("model should convert to response body", func(t *testing.T) {

		i := model.Ingredient{
			Name:   "someName",
			Amount: 1,
			Unit:   "someUnit",
		}
		resp := ingredientModelToResponse(i)

		assert.Equal(t, ingredientResponseBody{
			Name:   "someName",
			Unit:   "someUnit",
			Amount: 1,
		}, resp)
	})
}

func Test_ingredientRequestToModel(t *testing.T) {
	t.Run("request body should convert to model", func(t *testing.T) {

		i := ingredientRequestBody{
			Name:   "someName",
			Amount: 1,
			Unit:   "someUnit",
		}
		resp := ingredientRequestToModel(i)

		assert.Equal(t, model.Ingredient{
			Name:   "someName",
			Unit:   "someUnit",
			Amount: 1,
		}, resp)
	})
}

func Test_recipeToResponseModel(t *testing.T) {

	const id = 1

	r := model.Recipe{
		Id:     id,
		Author: "someauthor",
		Name:   "somename",
		Ingredients: []model.Ingredient{
			{

				Name:   "i",
				Amount: 1,
				Unit:   "unit",
			},
		},
		Directions:   []string{"do stuff"},
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}

	resp := recipeToResponseModel(r)

	assert.Equal(t, recipeResponseModel{
		Id:     id,
		Author: "someauthor",
		Name:   "somename",
		Ingredients: []ingredientResponseBody{
			{

				Name:   "i",
				Amount: 1,
				Unit:   "unit",
			},
		},
		Directions:   []string{"do stuff"},
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}, resp)
}

func Test_recipeRequestToModel(t *testing.T) {

	req := createRecipeRequestBody{
		Name: "somename",
		Ingredients: []ingredientRequestBody{
			{

				Name:   "i",
				Amount: 1,
				Unit:   "unit",
			},
		},
		Directions:   []string{"do stuff"},
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}

	recipe := recipeRequestToModel(req, "someuser")

	assert.Equal(t, model.Recipe{
		Id: 0,
		Author: "someuser",
		Name:   "somename",
		Ingredients: []model.Ingredient{
			{

				Name:   "i",
				Amount: 1,
				Unit:   "unit",
			},
		},
		Directions:   []string{"do stuff"},
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}, recipe)
}
