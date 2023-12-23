# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## auth-0.10.0 - 2023-12-23
#### Features
- **(authentication)** test change - (e731f3f) - TuftedDeer

- - -

## auth-0.9.0 - 2023-12-23
#### Features
- **(authentication)** test change - (6f9a810) - TuftedDeer

- - -

## auth-0.8.0 - 2023-12-23
#### Features
- **(authentication)** test change - (756c9e3) - TuftedDeer

- - -

## auth-0.7.0 - 2023-12-23
#### Features
- **(authentication)** test change - (6cb0494) - TuftedDeer

- - -

## auth-0.6.0 - 2023-12-23
#### Features
- **(authentication)** test change - (a20a409) - TuftedDeer

- - -

## auth-0.5.0 - 2023-12-23
#### Features
- **(authentication)** test change - (c09523e) - TuftedDeer

- - -

## auth-0.4.0 - 2023-12-23
#### Features
- **(authentication)** test change - (d9f5385) - TuftedDeer
#### Miscellaneous Chores
- **(version)** 0.1.0 - (8c494ec) - TuftedDeer

- - -

## auth-0.1.0 - 2023-12-23
#### Bug Fixes
- **(authentication)** listen on all interfaces - (74c7b09) - TuftedDeer
- **(recipe)** Add function for DeleteRecipe, modify DeleteRecipe for new RecipeId type, clean up imports a bit - (ad06026) - alex
#### Continuous Integration
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
#### Documentation
- add service readme files - (7392bea) - TuftedDeer
#### Features
- **(authentication)** test change - (385a490) - TuftedDeer
- **(authentication)** test change - (b8fd4f4) - TuftedDeer
- **(authentication)** test change - (5cb2d5a) - TuftedDeer
- **(authentication)** test change - (6c17056) - TuftedDeer
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
- breaking change - (842bd6a) - TuftedDeer
#### Miscellaneous Chores
- **(authentication)** fix "no such file" when starting container - (377d313) - TuftedDeer
- **(authentication)** go mod tidy - (67c8c21) - neki9072
- **(authentication)** Add Dockerfile for authentication - (7b1e3f9) - TuftedDeer
- **(version)** 0.1.0 - (6f3f28b) - Cog
- **(version)** 0.2.0 - (b5efe46) - Cog
- **(version)** 0.1.0 - (7df7597) - TuftedDeer
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## auth-0.1.0 - 2023-12-23
#### Bug Fixes
- **(authentication)** listen on all interfaces - (74c7b09) - TuftedDeer
- **(recipe)** Add function for DeleteRecipe, modify DeleteRecipe for new RecipeId type, clean up imports a bit - (ad06026) - alex
#### Continuous Integration
- **(authentication)** Use Ubuntu in Dockerfile build step, discard env variable name in build image workflow - (3e240f7) - alex
#### Documentation
- add service readme files - (7392bea) - TuftedDeer
#### Features
- **(authentication)** test change - (5cb2d5a) - TuftedDeer
- **(authentication)** test change - (6c17056) - TuftedDeer
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
- breaking change - (842bd6a) - TuftedDeer
#### Miscellaneous Chores
- **(authentication)** fix "no such file" when starting container - (377d313) - TuftedDeer
- **(authentication)** go mod tidy - (67c8c21) - neki9072
- **(authentication)** Add Dockerfile for authentication - (7b1e3f9) - TuftedDeer
- **(version)** 0.2.0 - (b5efe46) - Cog
- **(version)** 0.1.0 - (7df7597) - TuftedDeer
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## auth-0.2.0 - 2023-12-23
#### Features
- **(authentication)** test change - (6c17056) - TuftedDeer

- - -

## auth-0.1.0 - 2023-12-23
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
- breaking change - (842bd6a) - TuftedDeer
#### Miscellaneous Chores
- **(authentication)** fix "no such file" when starting container - (377d313) - TuftedDeer
- **(authentication)** go mod tidy - (67c8c21) - neki9072
- **(authentication)** Add Dockerfile for authentication - (7b1e3f9) - TuftedDeer
- **(version)** 0.4.0 - (c0948ee) - TuftedDeer
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

## auth-0.2.0 - 2023-12-23
#### Features
- breaking change - (842bd6a) - TuftedDeer

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).