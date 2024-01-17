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
        "timeEstimate",
        "difficulty",
        "feedsPeople",
        "directions",
        "author"
    )
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM
    "Recipe"
WHERE
    "recipeID" = $1;
