package recipes

import (
	"html/template"
	"net/http"

	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
)

func (ctrl *Controller) GetByAuthor(w http.ResponseWriter, r *http.Request) {

	// TODO: check if user exists, needs repository for accounts/users

	author := r.Context().Value("author").(string)

	recipes, err := ctrl.repo.GetAllByAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// template.Must(template.ParseGlob("templates/*.gohtml")) <--alex sagt
	response := fun.Map(recipes, recipeToResponseModel)

	tmplFile := "templates/displayRecipe.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	err = tmpl.Execute(w, response)
	if err != nil {
		panic(err)
	}

	//json_presenter.JsonPresenter(w, http.StatusOK, response)
}
