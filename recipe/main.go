package main

import (
	"context"
	"database/sql"
	"fmt"

	auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/environment"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/jwt_public_key"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"

	requestlogger "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request-logger"

	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
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

	jwtPublicKeyFile := environment.GetRequiredEnvVar(JwtPublicKeyEnvKey)
	pgConnString := environment.GetRequiredEnvVar(PostgresConnectionStringKey)

	fmt.Println("Hello from Recipe!")

	ctx := context.Background()

	// TODO: sslmode=disable should not be default, but it's too convenient atm
	db, err := sql.Open("pgx", pgConnString+"?sslmode=disable")
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

	authRpcTarget := environment.GetRequiredEnvVar(AuthRpcTarget)

	conn, err := grpc.Dial(authRpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	authRpcClient := auth_proto.NewAuthenticationClient(conn)

	recipeController := recipes.NewController(&repo, authRpcClient)

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
