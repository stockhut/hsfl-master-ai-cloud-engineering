// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"database/sql"
)

type Profiles struct {
	ProfileID      int64
	Username       string
	Password       string
	ProfilePicture []byte
	Bio            sql.NullString
	Friends        sql.NullInt64
	Weekplan       sql.NullInt64
}

type RecipeCollection struct {
	RecipeCollectionID   int64
	RecipeCollectionName string
	RecipeID             int64
	OwnerID              int64
	Date                 sql.NullString
	SubscriberID         sql.NullInt64
}

type Recipes struct {
	RecipeID      int64
	RecipeName    string
	RecipePicture []byte
	TimeEstimate  sql.NullInt64
	Difficulty    sql.NullString
	FeedsPeople   sql.NullInt64
	Ingredients   string
	Directions    string
	Author        string
}