-- name: GetAccount :one
SELECT *
FROM accounts
WHERE name = $1
LIMIT 1;

-- name: CreateAccount :execresult
INSERT INTO accounts (name, email, passwordhash)
VALUES ($1, $2, $3);

-- name: GetAccountByEmail :one
SELECT *
FROM accounts
WHERE email = $1
LIMIT 1;