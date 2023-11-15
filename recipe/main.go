package main

import (
	"context"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"

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

	jwtPrivateKeyFile, ok := os.LookupEnv(JwtPublicKeyEnvKey)
	if !ok {
		fmt.Printf("No %s configured\n", JwtPublicKeyEnvKey)
		os.Exit(1)
	}

	fmt.Println("Hello from Recipe!")

	ctx := context.Background()

	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, dll.Ddl); err != nil {
		panic(err)
	}

	queries := database.New(db)

	bytes, err := os.ReadFile(jwtPrivateKeyFile)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(bytes)
	public_key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	repo := recipes.New(queries)

	recipeController := recipes.NewController(&repo)

	authMiddleware := middleware.ValidateJwtMiddleware(public_key)

	logFlags := log.Ltime | log.Lmsgprefix | log.Lmicroseconds
	logger := log.New(os.Stdout, "", logFlags)
	logMw := requestlogger.New(logger)

	router := router.New(authMiddleware, logMw, recipeController)

	port := ":8081"

	logger.Printf("Listening on %s\n", port)
	err = http.ListenAndServe(port, router)
	panic(err)
}
