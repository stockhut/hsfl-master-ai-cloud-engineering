package recipes

import (
	auth_proto "github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/coalescing"
	"golang.org/x/sync/singleflight"
	"html/template"
)

type Controller struct {
	repo              RecipeRepository
	htmlTemplates     *template.Template
	authRpcClient     auth_proto.AuthenticationClient
	singleflightGroup coalescing.Coalescer
}

func NewController(repo RecipeRepository, authRpcClient auth_proto.AuthenticationClient, htmlTemplates *template.Template) *Controller {
	return &Controller{
		repo:              repo,
		authRpcClient:     authRpcClient,
		htmlTemplates:     htmlTemplates,
		singleflightGroup: &singleflight.Group{},
	}
}
