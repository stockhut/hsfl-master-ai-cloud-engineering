package reverse_proxy

import (
	"golang.org/x/exp/maps"
	"io"
	"net/http"
)

type Service struct {
	Name       string
	Route      string
	TargetHost string
}

// Forward the request to the given service, writing the service response to w
func Forward(w http.ResponseWriter, r *http.Request, service Service) error {

	r.Host = service.TargetHost

	newUrl := "http://" + service.TargetHost + r.URL.Path

	req, err := http.NewRequest(r.Method, newUrl, r.Body)
	if err != nil {
		return err
	}

	CopyHeaders(req.Header, r.Header)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	CopyHeaders(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func CopyHeaders(dest http.Header, src http.Header) {
	maps.Copy(dest, src)
}
