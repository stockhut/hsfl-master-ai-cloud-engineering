package recipes

import (
	"context"
	"database/sql"

	db "github.com/stockhut/hsfl-master-ai-cloud-engineering/common/db/generated"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
)

type RecipeRepository interface {
	GetAllByAuthor(author string) ([]model.Recipe, error)
	GetById(id model.RecipeId) (*model.Recipe, error)
	CreateRecipe(model.Recipe) (model.Recipe, error)
	DeleteRecipe(id model.RecipeId) error
}

type SqlcRepository struct {
	queries *db.Queries
}

func New(q *db.Queries) SqlcRepository {
	return SqlcRepository{queries: q}
}

func (repo *SqlcRepository) CreateRecipe(recipe model.Recipe) (model.Recipe, error) {

	params := db.CreateRecipeParams{
		RecipeName:   recipe.Name,
		TimeEstimate: sql.NullInt32{Int32: int32(recipe.TimeEstimate), Valid: true},
		Difficulty:   sql.NullString{String: recipe.Difficulty, Valid: true},
		Directions:   recipe.Directions[0],
		Author:       recipe.Author,
	}

	r, err := repo.queries.CreateRecipe(context.TODO(), params)

	if err != nil {
		return model.Recipe{}, err
	}

	ingredients := fun.Map(recipe.Ingredients, func(ingredient model.Ingredient) db.Ingredient {
		i, err := repo.queries.CreateIngredient(context.TODO(), db.CreateIngredientParams{
			RecipeID:         r.RecipeID,
			IngredientName:   ingredient.Name,
			IngredientAmount: int32(ingredient.Amount),
			IngredientUnit:   ingredient.Unit,
		})

		if err != nil {
			panic(err)
		}

		return i
	})

	return model.RecipeFromDatabaseModel(r, ingredients), nil
}

func (repo *SqlcRepository) GetAllByAuthor(author string) ([]model.Recipe, error) {

	r, err := repo.queries.ListRecipes(context.TODO(), author)
	if err != nil {
		return nil, err
	}

	recipes := fun.Map(r, func(recipe db.Recipe) model.Recipe {
		ingredients, _ := repo.queries.GetIngredientsByRecipe(context.TODO(), recipe.RecipeID)

		return model.RecipeFromDatabaseModel(recipe, ingredients)
	})

	return recipes, nil
}

func (repo *SqlcRepository) GetById(id model.RecipeId) (*model.Recipe, error) {
	recipe, err := repo.queries.GetRecipe(context.TODO(), int32(id))

	if err != nil {
		return nil, err
	}

	i, err := repo.queries.GetIngredientsByRecipe(context.TODO(), int32(id))

	if err != nil {
		return nil, err
	}

	result := model.RecipeFromDatabaseModel(recipe, i)
	return &result, nil
}

func (repo *SqlcRepository) DeleteRecipe(id model.RecipeId) error {
	return repo.queries.DeleteRecipe(context.TODO(), int32(id))
}
