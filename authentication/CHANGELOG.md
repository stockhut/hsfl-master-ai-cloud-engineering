# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## auth-0.5.0 - 2024-01-17
#### Bug Fixes
- **(authentication)** Stop logging jwts - (6c804a4) - TuftedDeer
#### Documentation
- update readmes - (797f1ee) - TuftedDeer
#### Features
- **(authentication)** Implement health check endpoint - (2a8a435) - TuftedDeer
#### Miscellaneous Chores
- **(docs)** fix table layout - (fcf0272) - Alexander Brandt
- **(docs)** add endpoints to readme - (8467ffc) - Alexander Brandt

- - -

## auth-0.4.0 - 2024-01-16
#### Bug Fixes
- **(authentication)** fix nil conversion error - (73834cd) - TuftedDeer
#### Documentation
- **(authentication)** fix proto generation instructions in readme - (2724ae5) - TuftedDeer
#### Features
- **(recipe)** Improve "get by author" performance under load through request coalescing - (5cd3112) - TuftedDeer
#### Miscellaneous Chores
- **(authentication)** use local common package version - (d9f21d2) - TuftedDeer

- - -

## auth-0.3.0 - 2024-01-07
#### Bug Fixes
- **(authentication)** jwt Cookie Path - (9f0d3ce) - neki9072
- **(authentication)** set jwt cookie expiration so it's not a session cookie - (2c8b06c) - TuftedDeer
#### Features
- add middleware capabilities to common router - (3526598) - TuftedDeer

- - -

## auth-0.2.0 - 2023-12-27
#### Features
- **(authentication)** use postgres db - (d78afaa) - TuftedDeer
- **(authentication)** use bcrypt to store passwords - (33453e8) - TuftedDeer
- **(authentication)** Handle account repo errors with ErrInternal gRPC status - (c2bdfec) - TuftedDeer
- **(authentication)** Add GRPC service - (abe1c70) - TuftedDeer
- **(recipe)** log rpc requests - (e8d54c2) - TuftedDeer
- **(recipe)** Verify account existence in GetByAuthor endpoint - (b8aaa67) - TuftedDeer
#### Miscellaneous Chores
- **(recipe)** Postgres deployment - (be579d0) - TuftedDeer
- generate protobuf code in Docker builds - (a59bace) - TuftedDeer
#### Refactoring
- **(authentication)** AccountRepository -> Repository - (56d65e0) - TuftedDeer
- **(authentication)** Use custom errors for account duplication check - (584f252) - TuftedDeer
- **(authentication)** remove repository subdirectory - (bfcc5c9) - TuftedDeer
- **(authentication)** add sqlc output because using sqlc with cgo in alpine container images is complicated - (0055fed) - TuftedDeer
- **(authentication)** move psql repo out of subdirectory - (cc6f54d) - TuftedDeer
#### Style
- **(authentication)** go fmt - (70254a8) - TuftedDeer
#### Tests
- **(authentication)** fix import path - (7d8f942) - TuftedDeer

- - -

## auth-0.1.0 - 2023-12-25
#### Bug Fixes
- **(authentication)** listen on all interfaces - (74c7b09) - TuftedDeer
- **(recipe)** Add function for DeleteRecipe, modify DeleteRecipe for new RecipeId type, clean up imports a bit - (ad06026) - alex
#### Continuous Integration
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
#### Documentation
- add service readme files - (7392bea) - TuftedDeer
#### Features
- **(authentication)** add required JWT_PRIVATE_KEY variable (#62) - (88a4a8a) - Fabian
- **(authentication)** add prefix to authentication routes, add reverse proxy setting - (44dfc34) - TuftedDeer
- **(authentication)** add jwt authentication middleware - (4720c8e) - TuftedDeer
- **(authentication)**  add jwt token to login - (37a3422) - Alexander Brandt
- **(authentication)** issue jwtToken at login - (4af860f) - neki9072
- **(authentication)** add account controller - (e651fe3) - Alexander Brandt
- **(authentication)** add accountRepository interface - (3a19db7) - TuftedDeer
- **(authentication)** create test for PostAccount - (9bde7f7) - Alexander Brandt
- **(authentication)** handle create account - (af9720c) - Alexander Brandt
- **(authentication)** temp database and handle login - (5f6c667) - neki9072
- **(authentication)** add basic http server - (387b0c7) - neki9072
- **(authentication/middleware)** Store only jwt claims, not the whole token - (a6b275e) - TuftedDeer
- **(common)** implement request logging - (d4a9398) - TuftedDeer
- **(recipes)** Add recipe service boilerplate and recipe creation endpoint - (f221688) - TuftedDeer
#### Miscellaneous Chores
- **(authentication)** fix "no such file" when starting container - (377d313) - TuftedDeer
- **(authentication)** go mod tidy - (67c8c21) - neki9072
- **(authentication)** Add Dockerfile for authentication - (7b1e3f9) - TuftedDeer
- tidy go.mod files - (b56f1d9) - TuftedDeer
#### Refactoring
- **(authentication)** authentiction refactoring - (804868b) - neki9072
- **(authentication/middleware)** use dedicated type for JWT Context key - (c4fb02d) - TuftedDeer
- remove unused code - (f974204) - TuftedDeer
#### Style
- **(authentication)** go fmt - (0a88cd2) - neki9072
- **(authentication/middleware)** go fmt - (243c69c) - TuftedDeer
- go fmt all the things - (0c2d052) - TuftedDeer
#### Tests
- **(authentication)** controller test - (d02a818) - neki9072
- **(authentication)** HandleCreateAccount Tests with Mocks - (b21a074) - neki9072
- **(authentication)** auth refactoring test with mocks +gitignore for mocks - (224d5b0) - neki9072
- **(authentication/middleware)** fix jwt authentication middleware construction invokation - (54edfda) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).