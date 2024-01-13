package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/environment"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/jwt_public_key"
	requestlogger "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request-logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"html/template"

	"log"

	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/pgxpool"
	dll "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db"
	database "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/middleware"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes"
)

const JwtPublicKeyEnvKey = "JWT_PUBLIC_KEY"
const PostgresConnectionStringKey = "PG_CONN_STRING"
const AuthRpcTarget = "AUTH_RPC_TARGET"

func main() {

	templates, err := template.ParseGlob("templates/**")

	if err != nil {
		log.Fatalf("Failed to parse html templates: %s\n", err)
	}

	jwtPublicKeyFile := environment.GetRequiredEnvVar(JwtPublicKeyEnvKey)
	pgConnString := environment.GetRequiredEnvVar(PostgresConnectionStringKey)

	fmt.Println("Hello from Recipe!")

	ctx := context.Background()

	dbPool, err := pgxpool.New(context.Background(), pgConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	_, err = dbPool.Exec(ctx, dll.Ddl)
	if err != nil {
		fmt.Printf("Failed to create tables: %s\n", err)
		return
	}

	queries := database.New(dbPool)

	publicKey, err := jwt_public_key.FromFile(jwtPublicKeyFile)
	if err != nil {
		fmt.Printf("Failed to load JWT public key: %s\n", err)
		return
	}

	repo := recipes.New(queries)

	authRpcTarget := environment.GetRequiredEnvVar(AuthRpcTarget)

	conn, err := grpc.Dial(authRpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	authRpcClient := auth_proto.NewAuthenticationClient(conn)

	recipeController := recipes.NewController(&repo, authRpcClient, templates)

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
