package json_presenter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonPresenter(t *testing.T) {

	type testStruct struct {
		Foo string `json:"foo"`
	}

	recorder := httptest.NewRecorder()

	Present(recorder, http.StatusTeapot, testStruct{Foo: "bar"})

	assert.Equal(t, http.StatusTeapot, recorder.Code)

	var body map[string]string
	err := json.NewDecoder(recorder.Result().Body).Decode(&body)

	assert.Nil(t, err)

	assert.Equal(t, "bar", body["foo"])
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
}
