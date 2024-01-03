package router

import (
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
)

type Router struct {
	router http.Handler
}

func New(
	authMiddleware func(http.HandlerFunc) http.HandlerFunc,
	logMiddleware func(http.HandlerFunc) http.HandlerFunc,
	recipeController *recipes.Controller,
) *Router {
	r := router.New(logMiddleware)

	r.POST("/api/v1/recipe", authMiddleware(recipeController.CreateRecipe))
	r.GET("/api/v1/recipe/by/self", authMiddleware(recipeController.GetBySelf))
	r.GET("/api/v1/recipe/by/:author", authMiddleware(recipeController.GetByAuthor))
	r.GET("/api/v1/recipe/:id", authMiddleware(recipeController.GetById))
	r.DELETE("/api/v1/recipe/:id", authMiddleware(recipeController.DeleteRecipe))
	r.GET("/health", recipeController.GetHealth)

	return &Router{r}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
