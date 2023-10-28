package recipes

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)
import "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/_mocks"

func TestDeleteRecipe(t *testing.T) {

	t.Run("should delete recipe", func(t *testing.T) {

		const testRecipeId = "testrecipe"
		const testUserName = "testuser"

		gomockController := gomock.NewController(t)

		mockRepo := mock_recipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().DeleteRecipe(model.RecipeId("testrecipe")).Return(nil).Times(1)

		controller := NewController(mockRepo)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/test", nil)

		claims := jwt.MapClaims{
			"name": testUserName,
		}
		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		ctx = context.WithValue(ctx, "id", "testrecipe")

		controller.DeleteRecipe(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusNoContent, w.Code)

	})

	t.Run("should return 500 INTERNAL SERVER ERROR when database write fails", func(t *testing.T) {

		const testRecipeId = "testrecipe"
		const testUserName = "testuser"

		gomockController := gomock.NewController(t)

		mockRepo := mock_recipes.NewMockRecipeRepository(gomockController)
		mockRepo.EXPECT().DeleteRecipe(model.RecipeId("testrecipe")).Return(errors.New("failed to delete recipe")).Times(1)

		controller := NewController(mockRepo)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/test", nil)

		claims := jwt.MapClaims{
			"name": testUserName,
		}
		ctx := context.WithValue(r.Context(), middleware.JwtContextKey, claims)
		ctx = context.WithValue(ctx, "id", testRecipeId)

		controller.DeleteRecipe(w, r.WithContext(ctx))
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

}
