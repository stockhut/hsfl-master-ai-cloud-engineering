package recipes

import (
	"net/http"
)

func (ctrl *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
