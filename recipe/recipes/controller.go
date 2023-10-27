package recipes

type Controller struct {
	repo RecipeRepository
}

func NewController(repo RecipeRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func mapx[T any, U any](ts []T, f func(T) U) []U {

	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}

	return us
}
