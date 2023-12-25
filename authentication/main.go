package main

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts/repository/postgresql"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/pwhash"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/environment"
	requestlogger "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/middleware/request-logger"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/accounts"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/api/router"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/jwt_util"
)

const JwtPrivateKeyEnvKey = "JWT_PRIVATE_KEY"
const PostgresConnectionStringKey = "PG_CONN_STRING"

func main() {

	jwtPrivateKeyFile := environment.GetRequiredEnvVar(JwtPrivateKeyEnvKey)
	pgConnString := environment.GetRequiredEnvVar(PostgresConnectionStringKey)

	tokenGeneratorConfig := jwt_util.JwtConfig{
		SignKey: jwtPrivateKeyFile,
	}
	tokenGenerator, err := jwt_util.NewJwtTokenGenerator(tokenGeneratorConfig)
	if err != nil {
		panic(err)
	}

	psqlRepo, err := postgresql.NewPsqlRepository(pgConnString)
	if err != nil {
		log.Fatalf("Failed to create psql repository: %s", err)
	}

	bcryptPwHasher := pwhash.BcryptPasswordHasher{}
	c := accounts.NewController(psqlRepo, *tokenGenerator, &bcryptPwHasher)

	fmt.Println("Hello from Auth!")

	logFlags := log.Ltime | log.Lmsgprefix | log.Lmicroseconds
	logger := log.New(os.Stdout, "", logFlags)
	logMw := requestlogger.New(logger)

	r := router.New(logMw, c)

	err = http.ListenAndServe(":8080", r)
	panic(err)
}
