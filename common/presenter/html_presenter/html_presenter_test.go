package html_presenter

import (
	"errors"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/_mocks/mock_html_presenter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPresent(t *testing.T) {

	t.Run("Html presenter", func(t *testing.T) {

		type Response struct {
			Value string
		}
		responseContent := Response{Value: "ok"}

		templateName := "test-template"

		t.Run("should render html", func(t *testing.T) {

			testTemplate, err := template.New(templateName).Parse("value: {{ .Value }}")

			assert.Nil(t, err)

			w := httptest.NewRecorder()
			Present(w, http.StatusTeapot, testTemplate, templateName, responseContent)

			body, err := io.ReadAll(w.Body)
			assert.Nil(t, err)

			assert.Contains(t, string(body), "value: ok")
			assert.Equal(t, http.StatusTeapot, w.Code)
			assert.Equal(t, "text/html", w.Header().Get("content-type"))
		})

		t.Run("should return 500 Internal Server Error if something goes wrong", func(t *testing.T) {

			mockController := gomock.NewController(t)
			mockTemplate := mock_html_presenter.NewMockHtmlTemplate(mockController)
			mockTemplate.EXPECT().ExecuteTemplate(gomock.Any(), templateName, responseContent).Return(errors.New("errmsg")).Times(1)
			mockTemplate.EXPECT().Name().Return(templateName).Times(1)

			w := httptest.NewRecorder()
			Present(w, http.StatusTeapot, mockTemplate, templateName, responseContent)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
