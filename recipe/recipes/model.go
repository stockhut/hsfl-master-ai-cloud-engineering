package recipes

type Ingredient struct {
	name   string
	amount int
	unit   string
}

type RecipeId string

type Recipe struct {
	id RecipeId
	// authors username
	author       string
	name         string
	ingredients  []Ingredient
	directions   []string
	timeEstimate int
	difficulty   string
	feedsPeople  int
}
