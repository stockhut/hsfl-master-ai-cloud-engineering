package recipes

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_GetHealth(t *testing.T) {

	recorder := httptest.NewRecorder()

	controller := NewController(nil, nil, nil)

	controller.GetHealth(recorder, nil)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
