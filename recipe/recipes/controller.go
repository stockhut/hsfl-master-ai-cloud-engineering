package recipes

import "html/template"

type Controller struct {
	repo          RecipeRepository
	htmlTemplates *template.Template
}

func NewController(repo RecipeRepository, htmlTemplates *template.Template) *Controller {
	return &Controller{
		repo:          repo,
		htmlTemplates: htmlTemplates,
	}
}
