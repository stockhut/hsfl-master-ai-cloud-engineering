# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## 0.8.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.8.0
### Global changes
#### Continuous Integration
- trigger docker build via gh cli if necessary - (56fdb5a) - TuftedDeer

- - -

## 0.7.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.7.0
### Global changes
#### Continuous Integration
- checkout repo before rebuild - (a41576f) - TuftedDeer

- - -

## 0.6.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.6.0
### Global changes
#### Continuous Integration
- run build action is neccessary - (9d16fb2) - TuftedDeer

- - -

## 0.5.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.5.0
### Global changes
#### Continuous Integration
- use auth-** in test and deliver - (e2353f6) - TuftedDeer
- hardcode test auth tag - (02dfbad) - TuftedDeer

- - -

## 0.4.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.4.0
### Global changes
#### Continuous Integration
- push cog commit - (a95c069) - TuftedDeer
#### Miscellaneous Chores
- **(version)** 0.1.0 - (8c494ec) - TuftedDeer

- - -

## 0.1.0 - 2023-12-23
### Package updates
- recipe bumped to recipe-0.1.0
- auth bumped to auth-0.1.0
- frontend bumped to frontend-0.1.0
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
- **(authentication)** deliver workflow names - (fcc69cd) - TuftedDeer
- **(authentication)** fix workflow path - (d5ccfac) - TuftedDeer
- **(authentication)** fix workflow path - (b4f237e) - TuftedDeer
- **(authentication)** better gh workflow file names - (8267906) - TuftedDeer
- **(authentication)** apparently GH workflows can't live in subdirectories - (ed4c252) - TuftedDeer
- **(authentication)** make test and docker workflows callable, add calling workflow for both - (91fbec7) - TuftedDeer
- **(authentication)** remove automatic run on test completion - (993bf14) - alex
- **(authentication)** install and run mockgen - (62c01df) - neki9072
- **(authentication)** remove test workflow dependency for now - (69db378) - alex
- **(authentication)** Add test worklow dependency - (1330ec6) - alex
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
- **(authentication)** run ci when yaml file changes - (57a1e62) - TuftedDeer
- **(authentication)** Add workflow to authentication service - (89b0ca8) - stockhut
- **(authentication/recipe)** callable workflows - (38fc88b) - TuftedDeer
- **(authentication/recipe)** callable workflows - (81a64d0) - TuftedDeer
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
- can i use tag OR path? - (f1b882d) - TuftedDeer
- push with tags - (7d103cf) - TuftedDeer
- push after cog - (d0112f6) - TuftedDeer
- git log - (27e3f6e) - TuftedDeer
- use pr head for version testing - (134a37c) - TuftedDeer
- run on versioning - (bcf7cd0) - TuftedDeer
- mark release as draft - (6368d0e) - TuftedDeer
- fix curl - (80a7274) - TuftedDeer
- fix gh repo variable - (1b0dc17) - TuftedDeer
- create tag using GH api - (15710c9) - TuftedDeer
- use cocogitto to create releases - (742c93d) - TuftedDeer
- fix on:path in recipe test and deploy - (0d7c992) - TuftedDeer
- move authentication stuff to subdirectory - (a538021) - TuftedDeer
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
- **(version)** 0.1.0 - (6f3f28b) - Cog
- **(version)** 0.2.0 - (b5efe46) - Cog
- **(version)** 0.1.0 - (7df7597) - TuftedDeer
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## 0.1.0 - 2023-12-23
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
- **(authentication)** deliver workflow names - (fcc69cd) - TuftedDeer
- **(authentication)** fix workflow path - (d5ccfac) - TuftedDeer
- **(authentication)** fix workflow path - (b4f237e) - TuftedDeer
- **(authentication)** better gh workflow file names - (8267906) - TuftedDeer
- **(authentication)** apparently GH workflows can't live in subdirectories - (ed4c252) - TuftedDeer
- **(authentication)** make test and docker workflows callable, add calling workflow for both - (91fbec7) - TuftedDeer
- **(authentication)** remove automatic run on test completion - (993bf14) - alex
- **(authentication)** install and run mockgen - (62c01df) - neki9072
- **(authentication)** remove test workflow dependency for now - (69db378) - alex
- **(authentication)** Add test worklow dependency - (1330ec6) - alex
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
- **(authentication)** run ci when yaml file changes - (57a1e62) - TuftedDeer
- **(authentication)** Add workflow to authentication service - (89b0ca8) - stockhut
- **(authentication/recipe)** callable workflows - (38fc88b) - TuftedDeer
- **(authentication/recipe)** callable workflows - (81a64d0) - TuftedDeer
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
- push after cog - (d0112f6) - TuftedDeer
- git log - (27e3f6e) - TuftedDeer
- use pr head for version testing - (134a37c) - TuftedDeer
- run on versioning - (bcf7cd0) - TuftedDeer
- mark release as draft - (6368d0e) - TuftedDeer
- fix curl - (80a7274) - TuftedDeer
- fix gh repo variable - (1b0dc17) - TuftedDeer
- create tag using GH api - (15710c9) - TuftedDeer
- use cocogitto to create releases - (742c93d) - TuftedDeer
- fix on:path in recipe test and deploy - (0d7c992) - TuftedDeer
- move authentication stuff to subdirectory - (a538021) - TuftedDeer
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
- **(version)** 0.2.0 - (b5efe46) - Cog
- **(version)** 0.1.0 - (7df7597) - TuftedDeer
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## 0.2.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.2.0
### Global changes
#### Continuous Integration
- push after cog - (d0112f6) - TuftedDeer
- git log - (27e3f6e) - TuftedDeer

- - -

## 0.1.0 - 2023-12-23
### Package updates
- frontend bumped to frontend-0.1.0
- recipe bumped to recipe-0.1.0
- auth bumped to auth-0.1.0
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
- **(authentication)** deliver workflow names - (fcc69cd) - TuftedDeer
- **(authentication)** fix workflow path - (d5ccfac) - TuftedDeer
- **(authentication)** fix workflow path - (b4f237e) - TuftedDeer
- **(authentication)** better gh workflow file names - (8267906) - TuftedDeer
- **(authentication)** apparently GH workflows can't live in subdirectories - (ed4c252) - TuftedDeer
- **(authentication)** make test and docker workflows callable, add calling workflow for both - (91fbec7) - TuftedDeer
- **(authentication)** remove automatic run on test completion - (993bf14) - alex
- **(authentication)** install and run mockgen - (62c01df) - neki9072
- **(authentication)** remove test workflow dependency for now - (69db378) - alex
- **(authentication)** Add test worklow dependency - (1330ec6) - alex
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
- **(authentication)** run ci when yaml file changes - (57a1e62) - TuftedDeer
- **(authentication)** Add workflow to authentication service - (89b0ca8) - stockhut
- **(authentication/recipe)** callable workflows - (38fc88b) - TuftedDeer
- **(authentication/recipe)** callable workflows - (81a64d0) - TuftedDeer
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
- use pr head for version testing - (134a37c) - TuftedDeer
- run on versioning - (bcf7cd0) - TuftedDeer
- mark release as draft - (6368d0e) - TuftedDeer
- fix curl - (80a7274) - TuftedDeer
- fix gh repo variable - (1b0dc17) - TuftedDeer
- create tag using GH api - (15710c9) - TuftedDeer
- use cocogitto to create releases - (742c93d) - TuftedDeer
- fix on:path in recipe test and deploy - (0d7c992) - TuftedDeer
- move authentication stuff to subdirectory - (a538021) - TuftedDeer
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
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## 0.4.0 - 2023-12-23
### Package updates
- auth bumped to auth-0.2.0
### Global changes
#### Continuous Integration
- **(authentication)** deliver workflow names - (fcc69cd) - TuftedDeer
- **(authentication)** fix workflow path - (d5ccfac) - TuftedDeer
- **(authentication)** fix workflow path - (b4f237e) - TuftedDeer
- **(authentication)** better gh workflow file names - (8267906) - TuftedDeer
- **(authentication)** apparently GH workflows can't live in subdirectories - (ed4c252) - TuftedDeer
- **(authentication)** make test and docker workflows callable, add calling workflow for both - (91fbec7) - TuftedDeer
- **(authentication/recipe)** callable workflows - (38fc88b) - TuftedDeer
- **(authentication/recipe)** callable workflows - (81a64d0) - TuftedDeer
- fix curl - (80a7274) - TuftedDeer
- fix gh repo variable - (1b0dc17) - TuftedDeer
- create tag using GH api - (15710c9) - TuftedDeer
- use cocogitto to create releases - (742c93d) - TuftedDeer
- fix on:path in recipe test and deploy - (0d7c992) - TuftedDeer
- move authentication stuff to subdirectory - (a538021) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).