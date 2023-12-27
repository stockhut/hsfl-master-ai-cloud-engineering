package recipes

import (
	"encoding/json"
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
		Directions:   "do stuff",
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
		Directions:   "do stuff",
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
		Directions:   "do stuff",
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}

	recipe := recipeRequestToModel(req, "someuser")

	assert.Equal(t, model.Recipe{
		Id:     0,
		Author: "someuser",
		Name:   "somename",
		Ingredients: []model.Ingredient{
			{

				Name:   "i",
				Amount: 1,
				Unit:   "unit",
			},
		},
		Directions:   "do stuff",
		TimeEstimate: 60,
		Difficulty:   "easy",
		FeedsPeople:  10,
	}, recipe)
}

func Test_createRecipeRequestBody_UnmarshalJSON(t *testing.T) {

	t.Run("createRecipeRequestBody UnmarshalJSON", func(t *testing.T) {

		t.Run("unmarshals ingredient attributes when they are string or []string", func(t *testing.T) {
			testCases := []struct {
				name                string
				input               string
				expectedIngredients []ingredientRequestBody
			}{
				{
					name: "ingredient-* as array",
					input: `
						{
						"ingredient-name":["first","second"],
						"ingredient-unit":["u1","u2"],
						"ingredient-amount":["1","2"]
						}`,
					expectedIngredients: []ingredientRequestBody{
						{
							Name:   "first",
							Unit:   "u1",
							Amount: 1,
						},
						{
							Name:   "second",
							Unit:   "u2",
							Amount: 2,
						},
					},
				},
				{
					name: "ingredient-* as single string",
					input: `
						{
						"ingredient-name":"first",
						"ingredient-unit":"u1",
						"ingredient-amount":"1"
						}`,
					expectedIngredients: []ingredientRequestBody{
						{
							Name:   "first",
							Unit:   "u1",
							Amount: 1,
						},
					},
				},
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					var obj createRecipeRequestBody

					err := json.Unmarshal([]byte(tc.input), &obj)

					assert.Nil(t, err)

					assert.Equal(t, tc.expectedIngredients, obj.Ingredients)
				})
			}

		})
	})
}
