----------PROFILE------------
-- name: GetProfile :one
SELECT
    *
FROM
    "Profile"
WHERE
    "profileID" = ?
LIMIT
    1;

-- name: ListProfiles :many
SELECT
    *
FROM
    "Profile"
ORDER BY
    "username";

-- name: CreateProfile :one
INSERT INTO
    "Profile" (
        "username",
        "password",
        "profilePicture",
        "bio"
    )
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: UpdateProfile :one
UPDATE
    "Profile"
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
    "Profile"
WHERE
    "profileID" = ?;

-----------INGREDIENT------------
-- name: GetIngredientsByRecipe :many
SELECT
    *
FROM 
    "Ingredient"
WHERE 
    "recipeID" = ?;

-- name: CreateIngredient :one
INSERT INTO
    "Ingredient" (
        "recipeID",
        "ingredientName",
        "ingredientAmount",
        "ingredientUnit"
    )
VALUES
    (?, ?, ?, ?) RETURNING *;


-----------RECIPE------------
-- name: GetRecipe :one
SELECT
    *
FROM
    "Recipe"
WHERE
    "recipeID" = ?
LIMIT
    1;

-- name: ListRecipes :many
SELECT
    *
FROM
    "Recipe"
WHERE
    "author" = ?
ORDER BY
    "recipeName";

-- name: CreateRecipe :one
INSERT INTO
    "Recipe" (
        "recipeName",
        "recipePicture",
        "timeEstimate",
        "difficulty",
        "feedsPeople",
        "directions",
        "author"
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateRecipe :one
UPDATE
    "Recipe"
set
    "recipeName" = ?,
    "recipePicture" = ?,
    "timeEstimate" = ?,
    "difficulty" = ?,
    "feedsPeople" = ?,
    "directions" = ?
WHERE
    "recipeID" = ? RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM
    "Recipe"
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