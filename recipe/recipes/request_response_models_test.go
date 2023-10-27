package recipes

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"testing"
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

	r := model.Recipe{
		Id:     "someid",
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
		Id:     "someid",
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
