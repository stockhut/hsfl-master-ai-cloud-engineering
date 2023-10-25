package recipes

import "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"

type RecipeRepository interface {
	//GetFromAuthor(author string) (Recipe, error)
	CreateRecipe(model.Recipe) (model.Recipe, error)
	//   DeleteRecipe(RecipeId) error
}

type InMemoryRecipeRepository struct {
	Recipes []model.Recipe
}

func (repo *InMemoryRecipeRepository) CreateRecipe(recipe model.Recipe) (model.Recipe, error) {

	recipe.Id = "TODO generate random id"
	repo.Recipes = append(repo.Recipes, recipe)

	return recipe, nil
}
