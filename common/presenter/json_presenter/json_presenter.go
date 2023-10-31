package json_presenter

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonPresenter[T any](w http.ResponseWriter, status int, content T) {

	body, err := json.Marshal(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(status)
	_, err = w.Write(body)

	if err != nil {
		log.Println(err)
	}
}
