package model

import "strconv"

type Ingredient struct {
	Name   string
	Amount int
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
	Directions   []string
	TimeEstimate int
	Difficulty   string
	FeedsPeople  int
}
