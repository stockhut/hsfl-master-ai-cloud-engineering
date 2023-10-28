package recipes

import "github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"

type RecipeRepository interface {
	GetAllByAuthor(author string) ([]model.Recipe, error)
	GetById(id model.RecipeId) (*model.Recipe, error)
	CreateRecipe(model.Recipe) (model.Recipe, error)
	DeleteRecipe(id model.RecipeId) error
}

type InMemoryRecipeRepository struct {
	Recipes []model.Recipe
}

func (repo *InMemoryRecipeRepository) CreateRecipe(recipe model.Recipe) (model.Recipe, error) {

	recipe.Id = "TODO generate random id"
	repo.Recipes = append(repo.Recipes, recipe)

	return recipe, nil
}

func (repo *InMemoryRecipeRepository) GetAllByAuthor(_ string) ([]model.Recipe, error) {
	panic("not implemented")
}

func (repo *InMemoryRecipeRepository) GetById(id model.RecipeId) (*model.Recipe, error) {
	panic("not implemented")
}
