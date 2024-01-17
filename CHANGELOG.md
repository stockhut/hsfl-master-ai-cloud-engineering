# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.8.0 - 2024-01-17
### Package updates
- recipe bumped to recipe-0.6.1
- auth bumped to auth-0.5.0
- common bumped to common-0.5.1
- frontend bumped to frontend-0.4.1
- loadbalancer bumped to loadbalancer-0.2.0
### Global changes
#### Miscellaneous Chores
- **(docs)** reference deployment readme in root readme - (ef529c1) - Alexander Brandt
- **(monitoring)** add readme - (7479165) - Alexander Brandt
- **(monitoring)** Display total pod RAM usage, display number of pods per image - (2d387ae) - TuftedDeer
- Add image build information to compose.yml - (747a545) - TuftedDeer

- - -

## 0.7.0 - 2024-01-17
### Package updates
- recipe bumped to recipe-0.6.0
- frontend bumped to frontend-0.4.0
### Global changes
#### Documentation
- update readmes - (797f1ee) - TuftedDeer

- - -

## 0.6.0 - 2024-01-16
### Package updates
- recipe bumped to recipe-0.5.0
- loadtest bumped to loadtest-0.3.0
- common bumped to common-0.5.0
- auth bumped to auth-0.4.0
### Global changes

- - -

## 0.5.0 - 2024-01-13
### Package updates
- loadtest bumped to loadtest-0.2.0
- common bumped to common-0.4.0
### Global changes
#### Continuous Integration
- **(loadtest)** fix workflow name - (1c64243) - Fabi
- **(loadtest)** Run tests - (865aa4b) - TuftedDeer
- **(web-service)** fetch tags when building web-service image - (adc6350) - TuftedDeer
- **(web-service)** don't ignore tags (trying to fix missing git tag when building image for new version) - (c95b91c) - TuftedDeer
#### Documentation
- **(deployment)** apply mainfests recusive (to include postgres) - (266c761) - TuftedDeer
- mention vm ocnfig in deployment readme - (16ed4c1) - TuftedDeer
- typo - (390000a) - Fabi
#### Miscellaneous Chores
- **(k8s-vm)** set cpu count to 2 - (d36690d) - TuftedDeer
- **(k8s-vm)** don't copy config into vm - (1de47bb) - TuftedDeer
- **(monitoring)** Add kubernetes dashboard (CPU/Mem per image, service network) - (4055124) - TuftedDeer
- add k3s vm config - (287ecf3) - TuftedDeer

- - -

## 0.4.0 - 2024-01-07
### Package updates
- frontend bumped to frontend-0.3.0
- auth bumped to auth-0.3.0
- common bumped to common-0.3.0
- recipe bumped to recipe-0.4.0
### Global changes
#### Bug Fixes
- **(monitoring)** use correct port for grafana service - (413adeb) - TuftedDeer
#### Continuous Integration
- **(release)** fix web-service build workflow path - (d2f1631) - Fabi
#### Documentation
- kubernetes + compose docs - (39559f1) - TuftedDeer
#### Miscellaneous Chores
- **(codecov)** informational mode for patches - (e0685d2) - TuftedDeer
- **(deployment)** add auth pg connection string, rpc connection config - (31f1536) - TuftedDeer
- **(monitoring)** update rate calculation and graph legends - (237ce46) - --local
- **(monitoring)** add custom dashboard to grafana - (8d28151) - --local
- **(monitoring)** add config files - (b93ba41) - --local
- update bruno requests - (3e4fafc) - TuftedDeer

- - -

## 0.3.0 - 2023-12-28
### Package updates
- recipe bumped to recipe-0.3.0
- reverseproxy bumped to reverseproxy-0.2.0
- frontend bumped to frontend-0.2.0
- common bumped to common-0.2.0
### Global changes
#### Continuous Integration
- **(recipe)** generate common mocks - (a9fcc13) - TuftedDeer
- rebuild reverse-proxy when new version is created - (454bb23) - TuftedDeer
#### Features
- **(common)** implement html presenter - (6ff0850) - TuftedDeer
- **(recipe)** make recipe directions a single string - (192e703) - TuftedDeer
- **(recipe)** allow more than one ingredient to be added via frontend - (b5c719a) - TuftedDeer
#### Miscellaneous Chores
- rename dev-db directory to dev-compose - (2d3d703) - TuftedDeer
- add frontend and reverse-proxy to compose - (2672894) - TuftedDeer

- - -

## 0.2.0 - 2023-12-27
### Package updates
- common bumped to common-0.1.0
- loadtest bumped to loadtest-0.1.0
- recipe bumped to recipe-0.2.0
- auth bumped to auth-0.2.0
- orchestration bumped to orchestration-0.1.0
- loadbalancer bumped to loadbalancer-0.1.0
- reverseproxy bumped to reverseproxy-0.1.0
### Global changes
#### Continuous Integration
- **(authentication)** generate mocks after protobuf - (9278d4c) - TuftedDeer
- **(authentication)** use go generate for protobuf generation - (954c10d) - TuftedDeer
- **(authentication)** install protoc-gen-go-grpc - (b943734) - TuftedDeer
- **(authentication)** fix protoc args - (cb622a8) - TuftedDeer
- **(authentication)** Generate protobuf code - (b127433) - TuftedDeer
- **(recipe)** Generate authentication mocks - (781416f) - TuftedDeer
- fix  image tags when no git tag exists, build images on all branches - (ba13356) - Fabi
- add remaining projects to cog.toml - (aca9105) - TuftedDeer
- add remaining projects to cog.toml - (dc30d2a) - TuftedDeer
- don't run auth and recipe test on every main push - (e35477f) - TuftedDeer
- build frontend image when creating a new release - (297feac) - TuftedDeer
- always run auth and recipe tests on:push:main when changed - (3c3061e) - TuftedDeer
- fix authentication rpc mock generation - (666a05e) - TuftedDeer
#### Documentation
- fix test badges, add version badges - (799ad1b) - TuftedDeer
#### Features
- **(authentication)** use bcrypt to store passwords - (33453e8) - TuftedDeer
- **(authentication)** Add GRPC service - (abe1c70) - TuftedDeer
- **(recipe)** Verify account existence in GetByAuthor endpoint - (b8aaa67) - TuftedDeer
#### Miscellaneous Chores
- **(bruno)** automatically extract jwt after login, mark jwt as secret - (9036fa3) - TuftedDeer
- **(codecov)** informational mode for patches - (d92ab74) - TuftedDeer
- **(recipe)** Postgres deployment - (be579d0) - TuftedDeer
- go mod tidy - (6b991ca) - TuftedDeer
- add recipe service to compose setup - (f854d40) - TuftedDeer
- put codecov in informational mode - (7c6e388) - TuftedDeer
- fix invalid codecov yaml - (86e8d8f) - TuftedDeer
- add auth service to compose setup - (a6ae6a1) - TuftedDeer
- explicitly set postgres user in compose - (8648802) - TuftedDeer
#### Refactoring
- **(authentication)** move psql repo out of subdirectory - (cc6f54d) - TuftedDeer

- - -

## 0.1.0 - 2023-12-25
### Package updates
- frontend bumped to frontend-0.1.0
- auth bumped to auth-0.1.0
- recipe bumped to recipe-0.1.0
### Global changes
#### Bug Fixes
- **(common)** create Tabe if not exists - (b98c2bc) - neki9072
- **(common/middleware/request-logger)** fix nil pointer deref when handler has not called WriteHeader - (f45bb2d) - TuftedDeer
- **(deployment)** remove broken volume configs - (b2e62c7) - TuftedDeer
- **(load-balancer)** fix race condition when passing healthyReplicas to balancing strategy - (fcd94ca) - TuftedDeer
- **(load-balancer)** Implement basic load balancer with round robin and least connections strategies - (8c0342e) - TuftedDeer
- **(load-balancer/round-robin)** fix race condition when updating index - (122bbd7) - TuftedDeer
- **(orchestration)** set container name and IP correctly - (b6d924d) - TuftedDeer
- **(orchestration)** stop orchestration from exiting - (4cdbad5) - TuftedDeer
- **(recipe)** Add function for DeleteRecipe, modify DeleteRecipe for new RecipeId type, clean up imports a bit - (ad06026) - alex
#### Continuous Integration
- **(authentication)** remove automatic run on test completion - (993bf14) - alex
- **(authentication)** install and run mockgen - (62c01df) - neki9072
- **(authentication)** remove test workflow dependency for now - (69db378) - alex
- **(authentication)** Add test worklow dependency - (1330ec6) - alex
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
- **(authentication)** run ci when yaml file changes - (57a1e62) - TuftedDeer
- **(authentication)** Add workflow to authentication service - (89b0ca8) - stockhut
- **(common)** add test workflow for common package - (0de546d) - TuftedDeer
- **(load-balancer)** remove automatic run on test completion - (7ffd5d3) - alex
- **(load-balancer)** remove test workflow dependency for now - (4b9ce3e) - alex
- **(load-balancer)** Add test worklow dependency - (bdc3040) - alex
- **(load-balancer)** Automate image build + push to GitHub registry - (0ca0189) - alex
- **(load-balancer)** Add Dockerfile - (ca3b68b) - alex
- **(recipe)** remove automatic run on test completion - (568191c) - alex
- **(recipe)** add workflow_dispatch to run workflow manually and change test worklow dependency logic - (fbd3d24) - alex
- **(recipe)** change test workflow name reference so this one actually triggers - (2f12114) - alex
- **(recipe)** Add workflow to recipe service - (dd7315e) - stockhut
- **(reverse-proxy)** remove automatic run on test completion - (110e402) - alex
- **(reverse-proxy)** remove test workflow dependency for now - (6d073db) - alex
- **(reverse-proxy)** Add test worklow dependency - (81899be) - alex
- **(reverse-proxy)** Discard env variable name in build image workflow - (a943122) - alex
- **(web-service)** remove automatic run on test completion - (f508da1) - alex
- **(web-service)** Add test worklow dependency - (986c913) - alex
- **(web-service)** Discard env variable name in build image workflow - (cb9ebcc) - alex
- automatically create SemVer tags and build container images - (35d8007) - Fabi
- use callable workflowsfor testing and container images - (9ac9e73) - Fabi
- add flags to reverseproxy and loadbalancer coverage uploads - (b7c5344) - TuftedDeer
- add test workflows for loadbalancer and reverse proxy - (449656d) - TuftedDeer
- mark codecov flags as carryforward, hoping to fix flaky coverage reports - (50526a3) - TuftedDeer
#### Documentation
- **(load-balancer)** add some docs about thread safety - (aaaf2f4) - TuftedDeer
- **(loadtest)** loadtest readme - (b9c3122) - TuftedDeer
- **(orchestration)** add readme - (313263d) - TuftedDeer
- add test status badge and goreportcard for reverse-proxy and load-balancer - (7ab1ab6) - alex
- add common test workflow badge - (4c33089) - TuftedDeer
- use correct goreportcard for common - (d351a85) - Fabian
- add codecov badge - (4979bcc) - alex
- add go report card badges - (e958141) - Fabian
- add test badges - (1f5e4e8) - Fabian
- add service readme files - (7392bea) - TuftedDeer
- update README - (1a5c1ec) - Alexander Brandt
#### Features
- **(authentication)** add prefix to authentication routes, add reverse proxy setting - (44dfc34) - TuftedDeer
- **(authentication)** add jwt authentication middleware - (4720c8e) - TuftedDeer
- **(authentication)** issue jwtToken at login - (4af860f) - neki9072
- **(authentication)** handle create account - (af9720c) - Alexander Brandt
- **(authentication)** add basic http server - (387b0c7) - neki9072
- **(bruno)** GetRecipeById Request - (6bf33df) - neki9072
- **(ci)** add codecov flags to each project - (947e67f) - TuftedDeer
- **(common)** implement json presenter function that handles json serialization and response writing - (078c150) - TuftedDeer
- **(common)** request logger recovers from panic in downstream handlers and responds with 500 - (2dd202c) - TuftedDeer
- **(common)** always color log output - (9bab32f) - TuftedDeer
- **(common)** colorize request logging output - (fd96358) - TuftedDeer
- **(common)** implement request logging - (d4a9398) - TuftedDeer
- **(common)** Add database querys - (e433466) - Alexander Brandt
- **(load-balancer)** implement ip hash strategy - (0328cc0) - TuftedDeer
- **(load-balancer)** initialize replica list with running container instances - (27a3eea) - TuftedDeer
- **(load-balancer)** turn LeastConnections.m into sync.RWMutex (was sync.Mutex) - (ebd6b8c) - TuftedDeer
- **(load-balancer)** make healthyLock sync.RWMutex (was sync.Mutex) - (5a01d23) - TuftedDeer
- **(load-balancer)** WIP: add container orchestration capabilities to load balancer - (de6d1bb) - TuftedDeer
- **(load-balancer)** Mark host as unhealthy when request forwarding fails - (be4e111) - TuftedDeer
- **(load-balancer)** Implement healthchecks - (b226931) - TuftedDeer
- **(loadtest)** let workers do http requests - (e943c87) - TuftedDeer
- **(loadtest)** implement ramped worker pool - (e87d32b) - TuftedDeer
- **(orchestration)** wait until image is pulled - (287ad4a) - TuftedDeer
- **(orchestration)** logging - (e6c11ab) - TuftedDeer
- **(orchestration)** pull images - (f373597) - TuftedDeer
- **(orchestration)** implement config file for orchestration - (aca3f80) - TuftedDeer
- **(orchestration)** automatically remove stopped containers - (9242b8f) - TuftedDeer
- **(orchestration)** move orchestration code from load-balancer to own project - (9b51c80) - TuftedDeer
- **(recipe)** Migrate from sqlite to PostgreSQL - (6917193) - TuftedDeer
- **(recipe)** add sqlite file location envrionment variable - (bdfa4e3) - TuftedDeer
- **(recipe)** sqlite database file instead of in-memory - (3d90177) - neki9072
- **(recipe)** add sqlite via sqlc - (4102772) - Alexander Brandt
- **(recipe)** add GetByAuthor endpoint - (8345f71) - TuftedDeer
- **(recipe)** ignore generated mock code files - (26c03b7) - TuftedDeer
- **(recipe)** CreateRecipe now returns lower case json, add test - (69b570e) - TuftedDeer
- **(recipes)** implement recipe list for the logged in user - (bde5834) - TuftedDeer
- **(recipes)** Add recipe service boilerplate and recipe creation endpoint - (f221688) - TuftedDeer
- **(reverse-proxy)** return 502 status when no service is found for requested path - (aca7236) - TuftedDeer
- **(reverse-proxy)** return 502 status when http forwarding fails - (d24f8d7) - TuftedDeer
- **(reverse-proxy)** implement config file loading - (dd54a93) - TuftedDeer
- **(reverse-proxy)** move request handling to proxy module, implement HttpHandler - (a73c7f8) - TuftedDeer
- **(reverse-proxy)** implement reverse proxy - (474b2ad) - TuftedDeer
- **(router)** add router implementation - (2718694) - neki9072
- **(web-service)** init static web server - (d74b8ba) - TuftedDeer
- init common go module - (a4520ba) - neki9072
#### Miscellaneous Chores
- **(bruno)** fix typo - (c82d9f0) - stockhut
- **(bruno)** add bruno collection - (3a3be6a) - TuftedDeer
- **(deployment)** Deployments yamls and persistentVolumes - (ac9c415) - neki9072
- **(recipe)** go mod tidy - (f2255ac) - neki9072
- **(recipe)** remove env variable for testing - (57559fe) - alex
- **(reverse-proxy)** Add Dockerfile for reverse-proxy - (e3eee84) - TuftedDeer
- **(reverse-proxy)** init - (f32dff2) - TuftedDeer
- add codecov attributes that might be missing - (1c2c58e) - Fabi
- add authentication deployment - (aaf716b) - TuftedDeer
- go mod tidy - (f45fa76) - TuftedDeer
- fix _mock subdirectories in .gitignore - (a599962) - TuftedDeer
- ignore editor, go and nix related files - (cf77a24) - TuftedDeer
- tidy go.mod files - (b56f1d9) - TuftedDeer
- add test running script - (d0c28fc) - TuftedDeer
#### Refactoring
- **(orchestration)** remove unused stuff - (5853af8) - TuftedDeer
- **(recipe)** move public key loading code to common - (2fb15b0) - TuftedDeer
- **(recipe)** move mapx function to common, rename to Map - (319e8b1) - TuftedDeer
- **(reverse-proxy)** move url to service matching to separate func - (41d1b94) - TuftedDeer
#### Style
- go fmt - (eb162a3) - TuftedDeer
#### Tests
- **(authentication)** auth refactoring test with mocks +gitignore for mocks - (224d5b0) - neki9072
- **(common)** implement more request logging tests - (7fedd0d) - TuftedDeer
- **(common)** test request logging middleware - (85126e0) - TuftedDeer
- **(common/fun)** test Map function - (f85547d) - TuftedDeer
- **(common/router)** test all implemented http methods - (277cef4) - TuftedDeer
- **(common/router)** add router tests - (51ef9dd) - TuftedDeer
- **(load-balancer/least-connections)** pass correct mutex type - (526d5e0) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).