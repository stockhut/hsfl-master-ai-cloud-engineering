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

	return &Router{router}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}