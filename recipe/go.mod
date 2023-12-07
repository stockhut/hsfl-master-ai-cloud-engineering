module github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

go 1.21.1

require (
	github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication v0.0.0-20231018220427-ee024c3b7b14
	github.com/stockhut/hsfl-master-ai-cloud-engineering/common v0.0.0-20231117115519-16be31557dde
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/jackc/pgx/v5 v5.5.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/mock v0.3.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/exp v0.0.0-20231206192017-f3f8817b8deb // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/stockhut/hsfl-master-ai-cloud-engineering/common => ../common

replace github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication => ../authentication
