package recipes

import "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"

type Controller struct {
	repo          RecipeRepository
	authRpcClient auth_proto.AuthenticationClient
}

func NewController(repo RecipeRepository, authRpcClient auth_proto.AuthenticationClient) *Controller {
	return &Controller{
		repo:          repo,
		authRpcClient: authRpcClient,
	}
}
