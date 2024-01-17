package model

import (
	"strconv"

	db "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

type Ingredient struct {
	Name   string
	Amount float64
	Unit   string
}

type RecipeId int64

func RecipeIdFromString(s string) (RecipeId, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	} else {
		return RecipeId(i), nil
	}
}

type Recipe struct {
	Id RecipeId
	// authors username
	Author       string
	Name         string
	Ingredients  []Ingredient
	Directions   string
	TimeEstimate int32
	Difficulty   string
	FeedsPeople  int32
}

func IngredientFromDatabaseModel(ingredient db.Ingredient) Ingredient {
	return Ingredient{
		Name:   ingredient.IngredientName,
		Amount: ingredient.IngredientAmount,
		Unit:   ingredient.IngredientUnit,
	}
}

func RecipeFromDatabaseModel(recipe db.Recipe, ingredients []db.Ingredient) Recipe {
	return Recipe{
		Id:           RecipeId(recipe.RecipeID),
		Author:       recipe.Author,
		Name:         recipe.RecipeName,
		Ingredients:  fun.Map(ingredients, IngredientFromDatabaseModel),
		Directions:   recipe.Directions,
		TimeEstimate: recipe.TimeEstimate,
		Difficulty:   recipe.Difficulty.String,
		FeedsPeople:  recipe.FeedsPeople,
	}
}
