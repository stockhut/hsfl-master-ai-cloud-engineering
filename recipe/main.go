package main

import (
	"context"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	dll "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db"
	database "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
)

func main() {

	fmt.Println("Hello from Recipe!")

	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, dll.Ddl); err != nil {
		panic(err)
	}

	queries := database.New(db)

	bytes, err := os.ReadFile("../authentication/jwt_public_key.key")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(bytes)
	public_key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	/* repo := recipes.InMemoryRecipeRepository{
		Recipes: make([]model.Recipe, 0),
	} */

	repo := recipes.New(queries)

	recipeController := recipes.NewController(&repo)

	authMiddleware := middleware.ValidateJwtMiddleware(public_key)

	router := router.New(authMiddleware, recipeController)

	err = http.ListenAndServe(":8081", router)
	panic(err)
}
