package recipes

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request_body"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)
import "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"

func TestCreateRecipe(t *testing.T) {

	t.Run("should create recipe", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mock_recipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().CreateRecipe(gomock.Any()).Return(model.Recipe{
			Id:           "some-id",
			Author:       "testuser",
			Name:         "",
			Ingredients:  nil,
			Directions:   nil,
			TimeEstimate: 0,
			Difficulty:   "",
			FeedsPeople:  0,
		}, nil).Times(1)

		controller := NewController(mockRepo)

		testBody := CreateRecipeRequestBody{
			Name: "my recipe",
			Ingredients: []ingredientRequestBody{
				{
					Name:   "rat",
					Amount: 1,
					Unit:   "pcs",
				},
			},
			Directions:   []string{"cook it"},
			TimeEstimate: 60,
			Difficulty:   "",
			FeedsPeople:  10,
		}

		w := httptest.NewRecorder()
		// body is set via context TODO: is this nice?
		r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

		claims := jwt.MapClaims{
			"name": "testuser",
		}
		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		ctx = context.WithValue(ctx, request_body.BodyMiddlewareContextKey, testBody)

		controller.CreateRecipe(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Equal(t, model.RecipeId("some-id"), responseBody.Id)
		assert.Equal(t, "testuser", responseBody.Author)
	})
}
