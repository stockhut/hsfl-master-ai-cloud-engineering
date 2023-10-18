package recipes

type RecipeRepository interface {
	//GetFromAuthor(author string) (Recipe, error)
	CreateRecipe(Recipe) (Recipe, error)
	//   DeleteRecipe(RecipeId) error
}

type InMemoryRecipeRepository struct {
	Recipes []Recipe
}

func (repo *InMemoryRecipeRepository) CreateRecipe(recipe Recipe) (Recipe, error) {

	recipe.id = "TODO generate random id"
	repo.Recipes = append(repo.Recipes, recipe)

	return recipe, nil
}
