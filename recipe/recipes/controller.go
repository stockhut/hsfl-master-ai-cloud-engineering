package recipes

import "html/template"

import "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"

type Controller struct {
	repo          RecipeRepository
	htmlTemplates *template.Template
	authRpcClient auth_proto.AuthenticationClient
}

func NewController(repo RecipeRepository, authRpcClient auth_proto.AuthenticationClient, htmlTemplates *template.Template) *Controller {
	return &Controller{
		repo:          repo,
		authRpcClient: authRpcClient,
		htmlTemplates: htmlTemplates,
	}
}
