package router

import (
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/router"
)

type Router struct {
	router http.Handler
}

func New(
	logMiddleware func(http.HandlerFunc) http.HandlerFunc,
	accountController *accounts.Controller,
) *Router {
	r := router.New(logMiddleware)

	r.POST("/api/v1/authentication/account", accountController.HandleCreateAccount)
	r.POST("/api/v1/authentication/login", accountController.HandleLogin)
	r.GET("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return &Router{r}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
