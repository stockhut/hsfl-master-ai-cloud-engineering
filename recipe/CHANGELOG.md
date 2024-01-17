# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## recipe-0.6.1 - 2024-01-17
#### Bug Fixes
- **(recipe)** return correct status code when recipe id is not found - (5fbc054) - TuftedDeer
- **(recipe)** fix time estimate and feeds people, remove unused sql code (without breaking things this time) - (e1c13b1) - TuftedDeer
- **(web-service)** fix input fields, required and remove Pfannkuchen - (d9d8516) - neki9072
#### Miscellaneous Chores
- **(docs)** add backticks to health endpoint - (8e048cf) - Alexander Brandt
- **(docs)** fix table layout - (fcf0272) - Alexander Brandt
- **(docs)** add endpoints to readme - (8467ffc) - Alexander Brandt
- go mod tidy - (ed62c63) - TuftedDeer

- - -

## recipe-0.6.0 - 2024-01-17
#### Bug Fixes
- **(recipe)** Display cooking instructions correctly/remove placeholder - (108d273) - TuftedDeer
#### Documentation
- update readmes - (797f1ee) - TuftedDeer
#### Features
- **(recipe)** tailwind css and html changes for recipe display - (840a2e4) - neki9072

- - -

## recipe-0.5.0 - 2024-01-16
#### Features
- **(recipe)** Use pgx connection pool to improve database connections - (26f2a5c) - TuftedDeer
- **(recipe)** Add profiling endpoint - (d06d815) - TuftedDeer
- **(recipe)** Improve "get by author" performance under load through request coalescing - (5cd3112) - TuftedDeer
#### Refactoring
- **(json-presenter)** rename function - (31dcaaf) - TuftedDeer

- - -

## recipe-0.4.0 - 2024-01-07
#### Features
- **(authentication)** make ingredient amount float - (a8707f6) - TuftedDeer
- add middleware capabilities to common router - (3526598) - TuftedDeer
#### Refactoring
- **(recipe)** remove third party pkg dependency - (46ec39e) - TuftedDeer

- - -

## recipe-0.3.0 - 2023-12-28
#### Bug Fixes
- **(recipe)** Fix no iteration over Directions - (9fac4f4) - neki9072
- **(recipe)** include html templates in docker image - (5966604) - TuftedDeer
#### Features
- **(recipe)** make recipe directions a single string - (192e703) - TuftedDeer
- **(recipe)** allow more than one ingredient to be added via frontend - (b5c719a) - TuftedDeer
#### Refactoring
- **(recipe)** parse html templates at start, use html presenter - (2b49395) - TuftedDeer
#### Style
- **(recipe)** go fmt - (debecf0) - TuftedDeer

- - -

## recipe-0.2.0 - 2023-12-27
#### Bug Fixes
- **(recipe)** self is not a path parameter - (7d26a1f) - TuftedDeer
#### Documentation
- **(recipe)** fix rpc description - (b138891) - TuftedDeer
#### Features
- **(recipe)** Verify account existence in GetByAuthor endpoint - (b8aaa67) - TuftedDeer
#### Miscellaneous Chores
- **(recipe)** go mod tidy - (650ccab) - TuftedDeer
- **(recipe)** Postgres deployment - (be579d0) - TuftedDeer
- **(recipe)** Bump authentication dependency, trying to fix go vet in ci - (c7e9522) - TuftedDeer
- generate protobuf code in Docker builds - (a59bace) - TuftedDeer
- go mod tidy - (6b991ca) - TuftedDeer
- go mod tidy - (b0b351e) - TuftedDeer

- - -

## recipe-0.1.0 - 2023-12-25
#### Bug Fixes
- **(recipe)** Return correct status code after recipe creation - (6376d10) - TuftedDeer
- **(recipe)** formatting via gofmt - (74f238b) - alex
- **(recipe)** Add function for DeleteRecipe, modify DeleteRecipe for new RecipeId type, clean up imports a bit - (ad06026) - alex
- **(recipe)** go mod tidy - (1a09e9c) - TuftedDeer
- use local version of common - (5d5fe4b) - TuftedDeer
#### Continuous Integration
- **(recipe)** go mod tidy :) - (71a3237) - alex
- **(recipe)** add replace for authentication - (f00430c) - alex
- **(recipe)** add @ to authentication version so it actually works :) - (2ee609f) - alex
- **(recipe)** try adding authentication directly - (f5a797a) - alex
- **(recipe)** change path to go.mod/go.sum - (9b022fc) - alex
- **(recipe)** remove apk install for gcc since it should already be included - (db66fae) - alex
- **(recipe)** change build image from alpine to what I think is ubuntu so the build step can use git :) - (17c8be9) - alex
#### Documentation
- add service readme files - (7392bea) - TuftedDeer
#### Features
- **(authentication/middleware)** Store only jwt claims, not the whole token - (a6b275e) - TuftedDeer
- **(common)** implement json presenter function that handles json serialization and response writing - (078c150) - TuftedDeer
- **(common)** implement request logging - (d4a9398) - TuftedDeer
- **(load-balancer)** implement ip hash strategy - (0328cc0) - TuftedDeer
- **(recipe)** Migrate from sqlite to PostgreSQL - (6917193) - TuftedDeer
- **(recipe)** add sqlite file location envrionment variable - (bdfa4e3) - TuftedDeer
- **(recipe)** log db opening and table creation errors - (58cb909) - TuftedDeer
- **(recipe)** add health check endpoint to recipe microservice - (dcba717) - Alexander Brandt
- **(recipe)** Add jwt public key environment config option - (eef1319) - TuftedDeer
- **(recipe)** create and show recipes in frontend - (aa67921) - neki9072
- **(recipe)** add delete recipe by id - (e9fa08b) - Alexander Brandt
- **(recipe)** recipe template parsing and displayRecipe Template - (f7be299) - neki9072
- **(recipe)** sqlite database file instead of in-memory - (3d90177) - neki9072
- **(recipe)** return database error instead of panicing - (999e1a2) - TuftedDeer
- **(recipe)** add sqlite via sqlc - (4102772) - Alexander Brandt
- **(recipe)** add DeleteRecipe route - (208a3c7) - TuftedDeer
- **(recipe)** implement DeleteRecipe endpoint - (b66e3b4) - TuftedDeer
- **(recipe)** add GetById endpoint - (51b5568) - TuftedDeer
- **(recipe)** add GetByAuthor endpoint - (8345f71) - TuftedDeer
- **(recipe)** CreateRecipe now returns HTTP 201-Created - (8b27896) - TuftedDeer
- **(recipe)** CreateRecipe now returns lower case json, add test - (69b570e) - TuftedDeer
- **(recipes)** implement recipe list for the logged in user - (bde5834) - TuftedDeer
- **(recipes)** read author name from jwt when creating new recie - (15d22d6) - TuftedDeer
- **(recipes)** Add recipe service boilerplate and recipe creation endpoint - (f221688) - TuftedDeer
- **(recipes/GetByAuthor)** send json or hypermedia, depending on the request header - (c979ba3) - TuftedDeer
#### Miscellaneous Chores
- **(recipe)** Remove sqlite stuff from Dockerfile - (d9c8267) - TuftedDeer
- **(recipe)** compile with cgo enabled - (b39b6f3) - TuftedDeer
- **(recipe)** go mod tidy - (f2255ac) - neki9072
- **(recipe)** Use alpine in dockerfile to fix image - (760ce61) - neki9072
- **(recipe)** Add recipe Dockerfile - (b3fcdf7) - TuftedDeer
- **(recipe)** fix go vet - (aa4a7ef) - TuftedDeer
- tidy go.mod files - (b56f1d9) - TuftedDeer
#### Refactoring
- **(authentication/middleware)** use dedicated type for JWT Context key - (c4fb02d) - TuftedDeer
- **(recipe)** router rename - (78a0230) - neki9072
- **(recipe)** move public key loading code to common - (2fb15b0) - TuftedDeer
- **(recipe)** variable naming - (de69fcd) - TuftedDeer
- **(recipe)** remove unused code - (b93ae1b) - TuftedDeer
- **(recipe)** remove unused in-memory repository - (468e7c4) - TuftedDeer
- **(recipe)** move model/database model conversion to functions - (c6ba0fe) - TuftedDeer
- **(recipe)** use test recipe id const - (9dccd67) - TuftedDeer
- **(recipe)** create converter function for recipe request bodies - (d904c8a) - TuftedDeer
- **(recipe)** move mapx function to common, rename to Map - (319e8b1) - TuftedDeer
- **(recipe)** move CreateRecipe handler to its own file - (d623941) - TuftedDeer
- **(recipe)** move request and response models to their own file - (1d60e89) - TuftedDeer
#### Style
- go fmt all the things - (0c2d052) - TuftedDeer
#### Tests
- **(recipe)** test health endpoint - (b1a0237) - TuftedDeer
- **(recipe)** fix recipe tests - (448004d) - TuftedDeer
- **(recipe)** test CreateRecipe behaviour when DB write fails - (5f71552) - TuftedDeer
- **(recipe)** test request and response model conversion - (8db8660) - TuftedDeer
- **(recipe)** make sure GetByAuthor handles users without recipes correct - (2fc02f1) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).