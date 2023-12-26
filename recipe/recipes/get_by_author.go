package recipes

import (
	"context"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/auth-proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"html/template"
	"log"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/htmx"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/presenter/json_presenter"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

func (ctrl *Controller) GetByAuthor(w http.ResponseWriter, r *http.Request) {

	author := r.Context().Value("author").(string)

	_, err := ctrl.authRpcClient.GetAccount(context.Background(), &auth_proto.GetAccountRequest{Name: author})
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

	recipes, err := ctrl.repo.GetAllByAuthor(author)
	if err != nil {
		log.Printf("Failed to get all by author: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// template.Must(template.ParseGlob("templates/*.gohtml")) <--alex sagt
	response := fun.Map(recipes, recipeToResponseModel)

	if htmx.IsHtmxRequest(r) {
		tmplFile := "templates/displayRecipesShort.html"
		tmpl, err := template.ParseFiles(tmplFile)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		err = tmpl.Execute(w, response)
		if err != nil {
			panic(err)
		}
	} else {
		json_presenter.JsonPresenter(w, http.StatusOK, response)
	}
}
