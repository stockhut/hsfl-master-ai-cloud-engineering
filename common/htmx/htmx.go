package htmx

import (
	"net/http"
)

// IsHtmxRequest inspects the request headers to identify whether it's an Ajax triggered by htmx
func IsHtmxRequest(r *http.Request) bool {
	// note the small x here!
	// while htmx uses HX-Request as key (https://htmx.org/reference/#request_headers),
	// net/http gives us canonical mime headers
	return r.Header.Get("Hx-Request") == "true"
}
