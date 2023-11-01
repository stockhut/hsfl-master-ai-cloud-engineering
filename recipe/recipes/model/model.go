package model

import (
	db "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

type Ingredient struct {
	Name   string
	Amount int
	Unit   string
}

type RecipeId int64

type Recipe struct {
	Id RecipeId
	// authors username
	Author       string
	Name         string
	Ingredients  []Ingredient
	Directions   []string
	TimeEstimate int
	Difficulty   string
	FeedsPeople  int
}

func IngredientFromDatabaseModel(ingredient db.Ingredient) Ingredient {
	return Ingredient{
		Name:   ingredient.IngredientName,
		Amount: int(ingredient.IngredientAmount),
		Unit:   ingredient.IngredientUnit,
	}
}

func RecipeFromDatabaseModel(recipe db.Recipe, ingredients []db.Ingredient) Recipe {
	return Recipe{
		Id:           RecipeId(recipe.RecipeID),
		Author:       recipe.Author,
		Name:         recipe.RecipeName,
		Ingredients:  fun.Map(ingredients, IngredientFromDatabaseModel),
		Directions:   []string{recipe.Directions},
		TimeEstimate: int(recipe.TimeEstimate.Int64),
		Difficulty:   recipe.Difficulty.String,
		FeedsPeople:  int(recipe.FeedsPeople.Int64),
	}
}
