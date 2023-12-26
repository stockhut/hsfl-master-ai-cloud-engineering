package html_presenter

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type HtmlTemplate interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
	Name() string
}

func Present[T any](w http.ResponseWriter, status int, tmpl HtmlTemplate, templateName string, content T) {

	out := make([]byte, 0)
	buffW := bytes.NewBuffer(out)

	err := tmpl.ExecuteTemplate(buffW, templateName, content)
	if err != nil {
		log.Printf("Failed to execute template '%s': %s\n", tmpl.Name(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(status)

	_, err = io.Copy(w, buffW)
	if err != nil {
		log.Printf("Failed to write response: %s\n", err)
	}
}
