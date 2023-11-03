package htmx

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

func TestIsHtmxRequest(t *testing.T) {

	t.Run("Returns the correct value for different requests", func(t *testing.T) {

		type testCase struct {
			header        http.Header
			isHtmxRequest bool
		}

		headerKey := textproto.CanonicalMIMEHeaderKey("HX-Request")

		testCases := []testCase{
			{
				header: map[string][]string{
					headerKey: {"true"},
				},
				isHtmxRequest: true,
			},
			{
				// this should not occur in real life, but better safe than sorry
				header: map[string][]string{
					headerKey: {"false"},
				},
				isHtmxRequest: false,
			},
			{
				header:        map[string][]string{},
				isHtmxRequest: false,
			},
		}

		for _, tc := range testCases {

			request := httptest.NewRequest(http.MethodGet, "/test", nil)
			request.Header = tc.header

			assert.Equal(t, tc.isHtmxRequest, IsHtmxRequest(request), "Test case failed: %v", tc)
		}
	})
}
