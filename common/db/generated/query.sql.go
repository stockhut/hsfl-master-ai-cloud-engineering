// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createIngredient = `-- name: CreateIngredient :one
INSERT INTO
    "Ingredient" (
        "recipeID",
        "ingredientName",
        "ingredientAmount",
        "ingredientUnit"
    )
VALUES
    ($1, $2, $3, $4) RETURNING "ingredientName", "ingredientAmount", "ingredientUnit", "recipeID"
`

type CreateIngredientParams struct {
	RecipeID         int32
	IngredientName   string
	IngredientAmount float64
	IngredientUnit   string
}

func (q *Queries) CreateIngredient(ctx context.Context, arg CreateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRow(ctx, createIngredient,
		arg.RecipeID,
		arg.IngredientName,
		arg.IngredientAmount,
		arg.IngredientUnit,
	)
	var i Ingredient
	err := row.Scan(
		&i.IngredientName,
		&i.IngredientAmount,
		&i.IngredientUnit,
		&i.RecipeID,
	)
	return i, err
}

const createRecipe = `-- name: CreateRecipe :one
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
    ($1, $2, $3, $4, $5, $6) RETURNING "recipeID", "recipeName", "timeEstimate", difficulty, "feedsPeople", directions, author
`

type CreateRecipeParams struct {
	RecipeName   string
	TimeEstimate int32
	Difficulty   pgtype.Text
	FeedsPeople  int32
	Directions   string
	Author       string
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRow(ctx, createRecipe,
		arg.RecipeName,
		arg.TimeEstimate,
		arg.Difficulty,
		arg.FeedsPeople,
		arg.Directions,
		arg.Author,
	)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.RecipeName,
		&i.TimeEstimate,
		&i.Difficulty,
		&i.FeedsPeople,
		&i.Directions,
		&i.Author,
	)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM
    "Recipe"
WHERE
    "recipeID" = $1
`

func (q *Queries) DeleteRecipe(ctx context.Context, recipeid int32) error {
	_, err := q.db.Exec(ctx, deleteRecipe, recipeid)
	return err
}

const getIngredientsByRecipe = `-- name: GetIngredientsByRecipe :many
SELECT
    "ingredientName", "ingredientAmount", "ingredientUnit", "recipeID"
FROM 
    "Ingredient"
WHERE 
    "recipeID" = $1
`

// ---------INGREDIENT------------
func (q *Queries) GetIngredientsByRecipe(ctx context.Context, recipeid int32) ([]Ingredient, error) {
	rows, err := q.db.Query(ctx, getIngredientsByRecipe, recipeid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ingredient
	for rows.Next() {
		var i Ingredient
		if err := rows.Scan(
			&i.IngredientName,
			&i.IngredientAmount,
			&i.IngredientUnit,
			&i.RecipeID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecipe = `-- name: GetRecipe :one
SELECT
    "recipeID", "recipeName", "timeEstimate", difficulty, "feedsPeople", directions, author
FROM
    "Recipe"
WHERE
    "recipeID" = $1
LIMIT
    1
`

// ---------RECIPE------------
func (q *Queries) GetRecipe(ctx context.Context, recipeid int32) (Recipe, error) {
	row := q.db.QueryRow(ctx, getRecipe, recipeid)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.RecipeName,
		&i.TimeEstimate,
		&i.Difficulty,
		&i.FeedsPeople,
		&i.Directions,
		&i.Author,
	)
	return i, err
}

const listRecipes = `-- name: ListRecipes :many
SELECT
    "recipeID", "recipeName", "timeEstimate", difficulty, "feedsPeople", directions, author
FROM
    "Recipe"
WHERE
    "author" = $1
ORDER BY
    "recipeName"
`

func (q *Queries) ListRecipes(ctx context.Context, author string) ([]Recipe, error) {
	rows, err := q.db.Query(ctx, listRecipes, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.RecipeID,
			&i.RecipeName,
			&i.TimeEstimate,
			&i.Difficulty,
			&i.FeedsPeople,
			&i.Directions,
			&i.Author,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
