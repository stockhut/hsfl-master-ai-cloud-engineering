package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"net/http"
	"os"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
)

func main() {

	fmt.Println("Hello from Recipe!")

	bytes, err := os.ReadFile("../authentication/jwt_public_key.key")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(bytes)
	public_key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	repo := recipes.InMemoryRecipeRepository{
		Recipes: make([]model.Recipe, 0),
	}

	recipeController := recipes.NewController(&repo)

	authMiddleware := middleware.ValidateJwtMiddleware(public_key)

	router := router.New(authMiddleware, recipeController)

	err = http.ListenAndServe(":8081", router)
	panic(err)
}
