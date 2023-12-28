package recipes

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	mock_auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/mock-auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	mockrecipes "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetBySelf(t *testing.T) {

	const testUserName = "testuser"
	claims := jwt.MapClaims{
		"name": testUserName,
	}

	t.Run("should return all recipes by a specific user", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().GetAllByAuthor(testUserName).Return([]model.Recipe{
			{
				Id:           1,
				Author:       testUserName,
				Name:         "",
				Ingredients:  nil,
				Directions:   "",
				TimeEstimate: 0,
				Difficulty:   "",
				FeedsPeople:  0,
			},
			{
				Id:           2,
				Author:       testUserName,
				Name:         "",
				Ingredients:  nil,
				Directions:   "",
				TimeEstimate: 0,
				Difficulty:   "",
				FeedsPeople:  0,
			},
		}, nil).Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		controller := NewController(mockRepo, mockAuthRpc, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		controller.GetBySelf(w, r.WithContext(ctx))

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody []recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Len(t, responseBody, 2)

		for _, recipe := range responseBody {
			assert.Equal(t, testUserName, recipe.Author)
		}
	})

	t.Run("should return 500 INTERNAL SERVER ERROR when database read fails", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.
			EXPECT().
			GetAllByAuthor(testUserName).
			Return(nil, errors.New("failed to read recipes")).
			Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		controller := NewController(mockRepo, mockAuthRpc, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/recipe", nil)

		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		controller.GetBySelf(w, r.WithContext(ctx))

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

	t.Run("should handle users without recipes correct", func(t *testing.T) {

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.
			EXPECT().
			GetAllByAuthor(testUserName).
			Return([]model.Recipe{}, nil).
			Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		controller := NewController(mockRepo, mockAuthRpc, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/recipe", nil)

		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		controller.GetBySelf(w, r.WithContext(ctx))

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody []recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Len(t, responseBody, 0)
	})
}
