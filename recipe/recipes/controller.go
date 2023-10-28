package recipes

type Controller struct {
	repo RecipeRepository
}

func NewController(repo RecipeRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}
