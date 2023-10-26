package recipes

import (
	"context"
	"encoding/json"
	"errors"
	mockrecipes "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetById(t *testing.T) {

	const testRecipeId = "test-recipe-id"

	t.Run("should return all recipes by a specific user", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().GetById(model.RecipeId(testRecipeId)).Return(&model.Recipe{
			Id:           testRecipeId,
			Author:       "testuser",
			Name:         "",
			Ingredients:  nil,
			Directions:   nil,
			TimeEstimate: 0,
			Difficulty:   "",
			FeedsPeople:  0,
		}, nil).Times(1)

		controller := NewController(mockRepo)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		ctx := context.WithValue(r.Context(), "id", testRecipeId)

		controller.GetById(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)
		assert.Equal(t, model.RecipeId(testRecipeId), responseBody.Id)

	})

	t.Run("should return 500 INTERNAL SERVER ERROR when database read fails", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.
			EXPECT().
			GetById(model.RecipeId(testRecipeId)).
			Return(nil, errors.New("failed to read recipes")).
			Times(1)

		controller := NewController(mockRepo)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		ctx := context.WithValue(r.Context(), "id", testRecipeId)

		controller.GetById(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

	t.Run("should return 404 NOT FOUND when there is no recipe with given id", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.
			EXPECT().
			GetById(model.RecipeId(testRecipeId)).
			Return(nil, nil).
			Times(1)

		controller := NewController(mockRepo)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		ctx := context.WithValue(r.Context(), "id", testRecipeId)

		controller.GetById(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusNotFound, w.Code)

	})
}
