package model

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
