package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/jwt_public_key"

	requestlogger "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request-logger"
	"log"

	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	dll "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db"
	database "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
)

const JwtPublicKeyEnvKey = "JWT_PUBLIC_KEY"

func main() {

	jwtPublicKeyFile, ok := os.LookupEnv(JwtPublicKeyEnvKey)
	if !ok {
		fmt.Printf("No %s configured\n", JwtPublicKeyEnvKey)
		os.Exit(1)
	}

	fmt.Println("Hello from Recipe!")

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, dll.Ddl); err != nil {
		fmt.Printf("Failed to create tables: %s\n", err)
		return
	}

	queries := database.New(db)

	publicKey, err := jwt_public_key.FromFile(jwtPublicKeyFile)
	if err != nil {
		fmt.Printf("Failed to load JWT public key: %s\n", err)
		return
	}

	repo := recipes.New(queries)

	recipeController := recipes.NewController(&repo)

	authMiddleware := middleware.ValidateJwtMiddleware(publicKey)

	logFlags := log.Ltime | log.Lmsgprefix | log.Lmicroseconds
	logger := log.New(os.Stdout, "", logFlags)
	logMw := requestlogger.New(logger)

	r := router.New(authMiddleware, logMw, recipeController)

	port := ":8081"

	logger.Printf("Listening on %s\n", port)
	err = http.ListenAndServe(port, r)
	panic(err)
}
