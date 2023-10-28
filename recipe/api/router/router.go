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
	recipeController *recipes.Controller,
) *Router {
	router := router.New()

	router.POST("/api/v1/recipe", authMiddleware(recipeController.CreateRecipe))
	router.GET("/api/v1/recipe/by/:author", authMiddleware(recipeController.GetByAuthor))
	router.GET("/api/v1/recipe/:id", authMiddleware(recipeController.GetById))
	router.DELETE("/api/v1/recipe/:id", authMiddleware(recipeController.DeleteRecipe))

	return &Router{router}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
