package router

import (
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
	"net/http"
)

type Router struct {
	router http.Handler
}

func New(
	authMiddleware func(http.HandlerFunc) http.HandlerFunc,
	logMiddleware func(http.HandlerFunc) http.HandlerFunc,
	recipeController *recipes.Controller,
) *Router {
	router := router.New()

	logAndAuth := fun.Apply(logMiddleware, authMiddleware)

	router.POST("/api/v1/recipe", logAndAuth(recipeController.CreateRecipe))
	router.GET("/api/v1/recipe/by/:author", logAndAuth(recipeController.GetByAuthor))
	router.GET("/api/v1/recipe/:id", logAndAuth(recipeController.GetById))
	router.DELETE("/api/v1/recipe/:id", logAndAuth(recipeController.DeleteRecipe))

	return &Router{router}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
