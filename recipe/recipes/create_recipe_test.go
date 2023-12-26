package recipes

import (
	"context"
	"encoding/json"
	"errors"
	mock_auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/mock-auth-proto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_recipes "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"
)

func TestCreateRecipe(t *testing.T) {

	const id = 1
	t.Run("should create recipe", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mock_recipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().CreateRecipe(gomock.Any()).Return(model.Recipe{
			Id:           id,
			Author:       "testuser",
			Name:         "",
			Ingredients:  nil,
			Directions:   nil,
			TimeEstimate: 0,
			Difficulty:   "",
			FeedsPeople:  0,
		}, nil).Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		controller := NewController(mockRepo, mockAuthRpc)

		testBody :=
			`{
				"name": "my recipe",
				"ingredients": [
					{
						"name": "rat",
						"unit": "pcs",
						"amount": 1
					}
				],
				"directions": [
					"cook it"
				],
				"time_estimate": 60,
				"difficulty": "easy",
				"feeds_people": 10
			}`

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/recipe", strings.NewReader(testBody))

		claims := jwt.MapClaims{
			"name": "testuser",
		}
		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)

		controller.CreateRecipe(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Equal(t, model.RecipeId(id), responseBody.Id)
		assert.Equal(t, "testuser", responseBody.Author)
	})

	t.Run("should return 500 INTERNAL SERVER ERROR when database write fails", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mock_recipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().CreateRecipe(gomock.Any()).Return(model.Recipe{}, errors.New("failed to save recipe")).Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		controller := NewController(mockRepo, mockAuthRpc)

		testBody :=
			`{
				"name": "my recipe",
				"ingredients": [
					{
						"name": "rat",
						"unit": "pcs",
						"amount": 1
					}
				],
				"directions": [
					"cook it"
				],
				"time_estimate": 60,
				"difficulty": "easy",
				"feeds_people": 10
			}`

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/recipe", strings.NewReader(testBody))

		claims := jwt.MapClaims{
			"name": "testuser",
		}
		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)

		controller.CreateRecipe(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
