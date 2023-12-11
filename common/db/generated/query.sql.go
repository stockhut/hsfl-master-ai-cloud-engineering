// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
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
	IngredientAmount int32
	IngredientUnit   string
}

func (q *Queries) CreateIngredient(ctx context.Context, arg CreateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRowContext(ctx, createIngredient,
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

const createProfile = `-- name: CreateProfile :one
INSERT INTO
    "Profile" (
        "username",
        "password",
        "profilePicture",
        "bio"
    )
VALUES
    ($1, $2, $3, $4) RETURNING "profileID", username, password, "profilePicture", bio, friends, weekplan
`

type CreateProfileParams struct {
	Username       string
	Password       string
	ProfilePicture []byte
	Bio            sql.NullString
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, createProfile,
		arg.Username,
		arg.Password,
		arg.ProfilePicture,
		arg.Bio,
	)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.Username,
		&i.Password,
		&i.ProfilePicture,
		&i.Bio,
		&i.Friends,
		&i.Weekplan,
	)
	return i, err
}

const createRecipe = `-- name: CreateRecipe :one
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
    ($1, $2, $3, $4, $5, $6, $7) RETURNING "recipeID", "recipeName", "recipePicture", "timeEstimate", difficulty, "feedsPeople", directions, author
`

type CreateRecipeParams struct {
	RecipeName    string
	RecipePicture []byte
	TimeEstimate  sql.NullInt32
	Difficulty    sql.NullString
	FeedsPeople   sql.NullInt32
	Directions    string
	Author        string
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.RecipeName,
		arg.RecipePicture,
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
		&i.RecipePicture,
		&i.TimeEstimate,
		&i.Difficulty,
		&i.FeedsPeople,
		&i.Directions,
		&i.Author,
	)
	return i, err
}

const createRecipeCollection = `-- name: CreateRecipeCollection :one
INSERT INTO
    "RecipeCollection" (
        "recipeCollectionName",
        "date"
    )
VALUES
    ($1, $2) RETURNING "recipeCollectionID", "recipeCollectionName", "recipeID", "ownerID", date, "subscriberID"
`

type CreateRecipeCollectionParams struct {
	RecipeCollectionName string
	Date                 sql.NullString
}

func (q *Queries) CreateRecipeCollection(ctx context.Context, arg CreateRecipeCollectionParams) (RecipeCollection, error) {
	row := q.db.QueryRowContext(ctx, createRecipeCollection, arg.RecipeCollectionName, arg.Date)
	var i RecipeCollection
	err := row.Scan(
		&i.RecipeCollectionID,
		&i.RecipeCollectionName,
		&i.RecipeID,
		&i.OwnerID,
		&i.Date,
		&i.SubscriberID,
	)
	return i, err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM
    "Profile"
WHERE
    "profileID" = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, profileid int32) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, profileid)
	return err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM
    "Recipe"
WHERE
    "recipeID" = $1
`

func (q *Queries) DeleteRecipe(ctx context.Context, recipeid int32) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, recipeid)
	return err
}

const deleteRecipeCollection = `-- name: DeleteRecipeCollection :exec
DELETE FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = $1
`

func (q *Queries) DeleteRecipeCollection(ctx context.Context, recipecollectionid int32) error {
	_, err := q.db.ExecContext(ctx, deleteRecipeCollection, recipecollectionid)
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
	rows, err := q.db.QueryContext(ctx, getIngredientsByRecipe, recipeid)
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
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfile = `-- name: GetProfile :one
SELECT
    "profileID", username, password, "profilePicture", bio, friends, weekplan
FROM
    "Profile"
WHERE
    "profileID" = $1
LIMIT
    1
`

// --------PROFILE------------
func (q *Queries) GetProfile(ctx context.Context, profileid int32) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, profileid)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.Username,
		&i.Password,
		&i.ProfilePicture,
		&i.Bio,
		&i.Friends,
		&i.Weekplan,
	)
	return i, err
}

const getRecipe = `-- name: GetRecipe :one
SELECT
    "recipeID", "recipeName", "recipePicture", "timeEstimate", difficulty, "feedsPeople", directions, author
FROM
    "Recipe"
WHERE
    "recipeID" = $1
LIMIT
    1
`

// ---------RECIPE------------
func (q *Queries) GetRecipe(ctx context.Context, recipeid int32) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, recipeid)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.RecipeName,
		&i.RecipePicture,
		&i.TimeEstimate,
		&i.Difficulty,
		&i.FeedsPeople,
		&i.Directions,
		&i.Author,
	)
	return i, err
}

const getRecipeCollection = `-- name: GetRecipeCollection :one
SELECT
    "recipeCollectionID", "recipeCollectionName", "recipeID", "ownerID", date, "subscriberID"
FROM
    "RecipeCollection"
WHERE
    "recipeCollectionID" = $1
LIMIT
    1
`

// -----RECIPECOLLECTION-------
func (q *Queries) GetRecipeCollection(ctx context.Context, recipecollectionid int32) (RecipeCollection, error) {
	row := q.db.QueryRowContext(ctx, getRecipeCollection, recipecollectionid)
	var i RecipeCollection
	err := row.Scan(
		&i.RecipeCollectionID,
		&i.RecipeCollectionName,
		&i.RecipeID,
		&i.OwnerID,
		&i.Date,
		&i.SubscriberID,
	)
	return i, err
}

const listProfiles = `-- name: ListProfiles :many
SELECT
    "profileID", username, password, "profilePicture", bio, friends, weekplan
FROM
    "Profile"
ORDER BY
    "username"
`

func (q *Queries) ListProfiles(ctx context.Context) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, listProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.ProfileID,
			&i.Username,
			&i.Password,
			&i.ProfilePicture,
			&i.Bio,
			&i.Friends,
			&i.Weekplan,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRecipeCollection = `-- name: ListRecipeCollection :many
SELECT
    "recipeCollectionID", "recipeCollectionName", "recipeID", "ownerID", date, "subscriberID"
FROM
    "RecipeCollection"
ORDER BY
    "recipeCollectionName"
`

func (q *Queries) ListRecipeCollection(ctx context.Context) ([]RecipeCollection, error) {
	rows, err := q.db.QueryContext(ctx, listRecipeCollection)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecipeCollection
	for rows.Next() {
		var i RecipeCollection
		if err := rows.Scan(
			&i.RecipeCollectionID,
			&i.RecipeCollectionName,
			&i.RecipeID,
			&i.OwnerID,
			&i.Date,
			&i.SubscriberID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRecipes = `-- name: ListRecipes :many
SELECT
    "recipeID", "recipeName", "recipePicture", "timeEstimate", difficulty, "feedsPeople", directions, author
FROM
    "Recipe"
WHERE
    "author" = $1
ORDER BY
    "recipeName"
`

func (q *Queries) ListRecipes(ctx context.Context, author string) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, listRecipes, author)
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
			&i.RecipePicture,
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
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProfile = `-- name: UpdateProfile :one
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
    "profileID" = $7 RETURNING "profileID", username, password, "profilePicture", bio, friends, weekplan
`

type UpdateProfileParams struct {
	Username       string
	Password       string
	ProfilePicture []byte
	Bio            sql.NullString
	Friends        sql.NullInt32
	Weekplan       sql.NullInt32
	ProfileID      int32
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, updateProfile,
		arg.Username,
		arg.Password,
		arg.ProfilePicture,
		arg.Bio,
		arg.Friends,
		arg.Weekplan,
		arg.ProfileID,
	)
	var i Profile
	err := row.Scan(
		&i.ProfileID,
		&i.Username,
		&i.Password,
		&i.ProfilePicture,
		&i.Bio,
		&i.Friends,
		&i.Weekplan,
	)
	return i, err
}

const updateRecipe = `-- name: UpdateRecipe :one
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
    "recipeID" = $7 RETURNING "recipeID", "recipeName", "recipePicture", "timeEstimate", difficulty, "feedsPeople", directions, author
`

type UpdateRecipeParams struct {
	RecipeName    string
	RecipePicture []byte
	TimeEstimate  sql.NullInt32
	Difficulty    sql.NullString
	FeedsPeople   sql.NullInt32
	Directions    string
	RecipeID      int32
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipe,
		arg.RecipeName,
		arg.RecipePicture,
		arg.TimeEstimate,
		arg.Difficulty,
		arg.FeedsPeople,
		arg.Directions,
		arg.RecipeID,
	)
	var i Recipe
	err := row.Scan(
		&i.RecipeID,
		&i.RecipeName,
		&i.RecipePicture,
		&i.TimeEstimate,
		&i.Difficulty,
		&i.FeedsPeople,
		&i.Directions,
		&i.Author,
	)
	return i, err
}

const updateRecipeCollection = `-- name: UpdateRecipeCollection :one
UPDATE
    "RecipeCollection"
set
    "recipeCollectionName" = $1,
    "date" = $2,
    "subscriberID" = $3
WHERE
    "recipeCollectionID" = $4 RETURNING "recipeCollectionID", "recipeCollectionName", "recipeID", "ownerID", date, "subscriberID"
`

type UpdateRecipeCollectionParams struct {
	RecipeCollectionName string
	Date                 sql.NullString
	SubscriberID         sql.NullInt32
	RecipeCollectionID   int32
}

func (q *Queries) UpdateRecipeCollection(ctx context.Context, arg UpdateRecipeCollectionParams) (RecipeCollection, error) {
	row := q.db.QueryRowContext(ctx, updateRecipeCollection,
		arg.RecipeCollectionName,
		arg.Date,
		arg.SubscriberID,
		arg.RecipeCollectionID,
	)
	var i RecipeCollection
	err := row.Scan(
		&i.RecipeCollectionID,
		&i.RecipeCollectionName,
		&i.RecipeID,
		&i.OwnerID,
		&i.Date,
		&i.SubscriberID,
	)
	return i, err
}
