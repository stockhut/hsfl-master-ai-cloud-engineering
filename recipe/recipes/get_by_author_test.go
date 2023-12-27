package recipes

import (
	"context"
	"encoding/json"
	"errors"
	mock_auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/_mocks/mock-auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	mockrecipes "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetByAuthor(t *testing.T) {

	t.Run("should return all recipes by a specific user", func(t *testing.T) {

		const testUserName = "testuser"

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().GetAllByAuthor(testUserName).Return([]model.Recipe{
			{
				Id:           1,
				Author:       testUserName,
				Name:         "",
				Ingredients:  nil,
				Directions:   nil,
				TimeEstimate: 0,
				Difficulty:   "",
				FeedsPeople:  0,
			},
			{
				Id:           2,
				Author:       testUserName,
				Name:         "",
				Ingredients:  nil,
				Directions:   nil,
				TimeEstimate: 0,
				Difficulty:   "",
				FeedsPeople:  0,
			},
		}, nil).Times(1)

		templates := template.Template{}
		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		mockAuthRpc.EXPECT().
			GetAccount(gomock.Any(), &auth_proto.GetAccountRequest{Name: testUserName}).
			Return(&auth_proto.GetAccountResponse{
				Name:  "testuser",
				Email: "testuser@example.org",
			}, nil).
			Times(1)

		controller := NewController(mockRepo, mockAuthRpc, &templates)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

		ctx := context.WithValue(r.Context(), "author", testUserName)

		controller.GetByAuthor(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody []recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Len(t, responseBody, 2)

		for _, recipe := range responseBody {
			assert.Equal(t, testUserName, recipe.Author)
		}
	})

	t.Run("should return 500 INTERNAL SERVER ERROR", func(t *testing.T) {

		t.Run(" when database read fails", func(t *testing.T) {

			const testUserName = "testuser"

			gomockController := gomock.NewController(t)

			mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
			mockRepo.
				EXPECT().
				GetAllByAuthor(testUserName).
				Return(nil, errors.New("failed to read recipes")).
				Times(1)

			mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
			mockAuthRpc.EXPECT().
				GetAccount(gomock.Any(), &auth_proto.GetAccountRequest{Name: testUserName}).
				Return(&auth_proto.GetAccountResponse{
					Name:  "testuser",
					Email: "testuser@example.org",
				}, nil).
				Times(1)

			controller := NewController(mockRepo, mockAuthRpc, nil)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

			ctx := context.WithValue(r.Context(), "author", testUserName)

			controller.GetByAuthor(w, r.WithContext(ctx))
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run(" when auth rpc returns a fatal error", func(t *testing.T) {

			const testUserName = "testuser"

			gomockController := gomock.NewController(t)

			mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)

			mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
			mockAuthRpc.EXPECT().
				GetAccount(gomock.Any(), &auth_proto.GetAccountRequest{Name: testUserName}).
				Return(&auth_proto.GetAccountResponse{}, errors.New("something really bad happened")).
				Times(1)

		templates := template.Template{}
			controller := NewController(mockRepo, mockAuthRpc, &templates)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

			ctx := context.WithValue(r.Context(), "author", testUserName)

			controller.GetByAuthor(w, r.WithContext(ctx))
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("should handle users without recipes correct", func(t *testing.T) {

		const testUserName = "testuser"

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)
		mockRepo.
			EXPECT().
			GetAllByAuthor(testUserName).
			Return([]model.Recipe{}, nil).
			Times(1)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		mockAuthRpc.EXPECT().
			GetAccount(gomock.Any(), &auth_proto.GetAccountRequest{Name: testUserName}).
			Return(&auth_proto.GetAccountResponse{
				Name:  "testuser",
				Email: "testuser@example.org",
			}, nil).
			Times(1)

		templates := template.Template{}
		controller := NewController(mockRepo, mockAuthRpc, &templates)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

		ctx := context.WithValue(r.Context(), "author", testUserName)

		controller.GetByAuthor(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody []recipeResponseModel
		err := json.NewDecoder(w.Body).Decode(&responseBody)

		assert.Nil(t, err)

		assert.Len(t, responseBody, 0)
	})

	t.Run("returns 404 when user does not exist", func(t *testing.T) {

		const testUserName = "testuser"

		gomockController := gomock.NewController(t)

		mockRepo := mockrecipes.NewMockRecipeRepository(gomockController)

		mockAuthRpc := mock_auth_proto.NewMockAuthenticationClient(gomockController)
		mockAuthRpc.EXPECT().
			GetAccount(gomock.Any(), &auth_proto.GetAccountRequest{Name: testUserName}).
			Return(&auth_proto.GetAccountResponse{}, auth_proto.ErrAccountNotFound).
			Times(1)

		templates := template.Template{}
		controller := NewController(mockRepo, mockAuthRpc, &templates)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/recipe", nil)

		ctx := context.WithValue(r.Context(), "author", testUserName)

		controller.GetByAuthor(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
