module github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe

go 1.21.1

require (
	github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication v0.0.0-20231018220427-ee024c3b7b14
	github.com/stockhut/hsfl-master-ai-cloud-engineering/common v0.0.0-20231117115519-16be31557dde
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/stretchr/testify v1.8.4
	go.uber.org/mock v0.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/stockhut/hsfl-master-ai-cloud-engineering/common => ../common

replace github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication => ../authentication
