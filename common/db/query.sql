----------PROFILES------------
-- name: GetProfile :one
SELECT
    *
FROM
    "Profiles"
WHERE
    "profileID" = ?
LIMIT
    1;

-- name: ListProfiles :many
SELECT
    *
FROM
    "Profiles"
ORDER BY
    "username";

-- name: CreateProfile :one
INSERT INTO
    "Profiles" (
        "username",
        "password",
        "profilePicture",
        "bio"
    )
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: UpdateProfile :one
UPDATE
    "Profiles"
set
    "username" = ?,
    "password" = ?,
    "profilePicture" = ?,
    "bio" = ?,
    "friends" = ?,
    "weekplan" = ?
WHERE
    "profileID" = ? RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM
    "Profiles"
WHERE
    "profileID" = ?;

-----------RECIPE------------
-- name: GetRecipe :one
SELECT
    *
FROM
    "Recipes"
WHERE
    "recipeID" = ?
LIMIT
    1;

-- name: ListRecipes :many
SELECT
    *
FROM
    "Recipes"
ORDER BY
    "recipeName";

-- name: CreateRecipe :one
INSERT INTO
    "Recipes" (
        "recipeName",
        "recipePicture",
        "timeEstimate",
        "difficulty",
        "feedsPeople",
        "ingredients",
        "directions"
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateRecipe :one
UPDATE
    "Recipes"
set
    "recipeName" = ?,
    "recipePicture" = ?,
    "timeEstimate" = ?,
    "difficulty" = ?,
    "feedsPeople" = ?,
    "ingredients" = ?,
    "directions" = ?
WHERE
    "recipeID" = ? RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM
    "Recipes"
WHERE
    "recipeID" = ?;

-------RECIPECOLLECTION-------
-- name: GetRecipeCollection :one
SELECT
    *
FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = ?
LIMIT
    1;

-- name: ListRecipeCollection :many
SELECT
    *
FROM
    "RecipeCollection"
ORDER BY
    "recipeCollectionName";

-- name: CreateRecipeCollection :one
INSERT INTO
    "RecipeCollection" (
        "recipeCollectionName",
        "date"
    )
VALUES
    (?, ?) RETURNING *;

-- name: UpdateRecipeCollection :one
UPDATE
    "RecipeCollection"
set
    "recipeCollectionName" = ?,
    "date" = ?,
    "subscriberID" = ?
WHERE
    "recipeCollectionID" = ? RETURNING *;

-- name: DeleteRecipeCollection :exec
DELETE FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = ?;