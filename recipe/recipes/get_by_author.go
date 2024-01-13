package recipes

import (
	"context"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/html_presenter"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/recipe/recipes/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

func (ctrl *Controller) GetByAuthor(w http.ResponseWriter, r *http.Request) {

	author := r.Context().Value("author").(string)

	_, err, _ := ctrl.singleflightGroup.Do("get-author-rpc "+author, func() (interface{}, error) {
		return ctrl.authRpcClient.GetAccount(context.Background(), &auth_proto.GetAccountRequest{Name: author})
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Printf("Failed to call GetAccount over RPC: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	recipes, err, _ := ctrl.singleflightGroup.Do("get-all-by-author "+author, func() (interface{}, error) {
		return ctrl.repo.GetAllByAuthor(author)
	})
	if err != nil {
		log.Printf("Failed to get all by author: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// template.Must(template.ParseGlob("templates/*.gohtml")) <--alex sagt
	response := fun.Map(recipes.([]model.Recipe), recipeToResponseModel)
	if htmx.IsHtmxRequest(r) {
		html_presenter.Present(w, http.StatusOK, ctrl.htmlTemplates, "displayRecipesShort.html", response)
	} else {
		json_presenter.JsonPresenter(w, http.StatusOK, response)
	}
}
