----------PROFILE------------
-- name: GetProfile :one
SELECT
    *
FROM
    "Profile"
WHERE
    "profileID" = $1
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
    ($1, $2, $3, $4) RETURNING *;

-- name: UpdateProfile :one
UPDATE
    "Profile"
set
    "username" = $1,
    "password" = $2,
    "profilePicture" = $3,
    "bio" = $4,
    "friends" = $5,
    "weekplan" = $6
WHERE
    "profileID" = $7 RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM
    "Profile"
WHERE
    "profileID" = $1;

-----------INGREDIENT------------
-- name: GetIngredientsByRecipe :many
SELECT
    *
FROM 
    "Ingredient"
WHERE 
    "recipeID" = $1;

-- name: CreateIngredient :one
INSERT INTO
    "Ingredient" (
        "recipeID",
        "ingredientName",
        "ingredientAmount",
        "ingredientUnit"
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;


-----------RECIPE------------
-- name: GetRecipe :one
SELECT
    *
FROM
    "Recipe"
WHERE
    "recipeID" = $1
LIMIT
    1;

-- name: ListRecipes :many
SELECT
    *
FROM
    "Recipe"
WHERE
    "author" = $1
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
    ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdateRecipe :one
UPDATE
    "Recipe"
set
    "recipeName" = $1,
    "recipePicture" = $2,
    "timeEstimate" = $3,
    "difficulty" = $4,
    "feedsPeople" = $5,
    "directions" = $6
WHERE
    "recipeID" = $7 RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM
    "Recipe"
WHERE
    "recipeID" = $1;

-------RECIPECOLLECTION-------
-- name: GetRecipeCollection :one
SELECT
    *
FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = $1
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
    ($1, $2) RETURNING *;

-- name: UpdateRecipeCollection :one
UPDATE
    "RecipeCollection"
set
    "recipeCollectionName" = $1,
    "date" = $2,
    "subscriberID" = $3
WHERE
    "recipeCollectionID" = $4 RETURNING *;

-- name: DeleteRecipeCollection :exec
DELETE FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = $1;