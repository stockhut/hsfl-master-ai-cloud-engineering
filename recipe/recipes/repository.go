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

type InMemoryRecipeRepository struct {
	Recipes []model.Recipe
}

func (repo *InMemoryRecipeRepository) CreateRecipe(recipe model.Recipe) (model.Recipe, error) {

	recipe.Id = 0 //"TODO generate random id"
	repo.Recipes = append(repo.Recipes, recipe)

	return recipe, nil
}

func (repo *SqlcRepository) CreateRecipe(recipe model.Recipe) (model.Recipe, error) {

	params := db.CreateRecipeParams{
		RecipeName:   recipe.Name,
		TimeEstimate: sql.NullInt64{Int64: int64(recipe.TimeEstimate), Valid: true},
		Difficulty:   sql.NullString{String: recipe.Difficulty, Valid: true},
		Directions:   recipe.Directions[0],
		Author:       recipe.Author,
	}

	r, err := repo.queries.CreateRecipe(context.TODO(), params)

	if err != nil {
		panic(err)
	}

	ingredients := fun.Map(recipe.Ingredients, func(ingredient model.Ingredient) model.Ingredient {
		i, err := repo.queries.CreateIngredient(context.TODO(), db.CreateIngredientParams{
			RecipeID:         r.RecipeID,
			IngredientName:   ingredient.Name,
			IngredientAmount: int64(ingredient.Amount),
			IngredientUnit:   ingredient.Unit,
		})

		if err != nil {
			panic(err)
		}

		return model.Ingredient{
			Name:   i.IngredientName,
			Amount: ingredient.Amount,
			Unit:   i.IngredientUnit,
		}
	})

	return model.Recipe{
		Id:           model.RecipeId(r.RecipeID),
		Author:       r.Author,
		Name:         r.RecipeName,
		Ingredients:  ingredients,
		Directions:   []string{r.Directions},
		TimeEstimate: int(r.TimeEstimate.Int64),
		Difficulty:   r.Difficulty.String,
		FeedsPeople:  int(r.FeedsPeople.Int64),
	}, nil
}

func (repo *InMemoryRecipeRepository) GetAllByAuthor(_ string) ([]model.Recipe, error) {
	panic("not implemented")
}

func (repo *SqlcRepository) GetAllByAuthor(author string) ([]model.Recipe, error) {

	r, err := repo.queries.ListRecipes(context.TODO(), author)

	recipes := fun.Map(r, func(recipe db.Recipe) model.Recipe {
		i, _ := repo.queries.GetIngredientsByRecipe(context.TODO(), recipe.RecipeID)

		ingredients := fun.Map(i, func(ingredient db.Ingredient) model.Ingredient {
			return model.Ingredient{
				Name:   ingredient.IngredientName,
				Amount: int(ingredient.IngredientAmount),
				Unit:   ingredient.IngredientUnit,
			}
		})

		return model.Recipe{
			Id:           model.RecipeId(recipe.RecipeID),
			Author:       recipe.Author,
			Name:         recipe.RecipeName,
			Ingredients:  ingredients,
			Directions:   []string{recipe.Directions},
			TimeEstimate: int(recipe.TimeEstimate.Int64),
			Difficulty:   recipe.Difficulty.String,
			FeedsPeople:  int(recipe.FeedsPeople.Int64),
		}
	})

	return recipes, err
}

func (repo *InMemoryRecipeRepository) GetById(id model.RecipeId) (*model.Recipe, error) {
	panic("not implemented")
}

func (repo *SqlcRepository) GetById(id model.RecipeId) (*model.Recipe, error) {
	recipe, err := repo.queries.GetRecipe(context.TODO(), int64(id))

	if err != nil {
		panic(err)
	}

	i, err := repo.queries.GetIngredientsByRecipe(context.TODO(), int64(id))

	if err != nil {
		panic(err)
	}

	ingredients := fun.Map(i, func(ingredient db.Ingredient) model.Ingredient {
		return model.Ingredient{
			Name:   ingredient.IngredientName,
			Amount: int(ingredient.IngredientAmount),
			Unit:   ingredient.IngredientUnit,
		}
	})
	return &model.Recipe{
		Id:           model.RecipeId(id),
		Author:       recipe.Author,
		Name:         recipe.RecipeName,
		Ingredients:  ingredients,
		Directions:   []string{recipe.Directions},
		TimeEstimate: int(recipe.TimeEstimate.Int64),
		Difficulty:   recipe.Difficulty.String,
		FeedsPeople:  int(recipe.FeedsPeople.Int64),
	}, err
}

func (repo *InMemoryRecipeRepository) DeleteRecipe(id model.RecipeId) error {
	panic("not implemented")
}

func (repo *SqlcRepository) DeleteRecipe(id model.RecipeId) error {
	panic("not implemented")
}
