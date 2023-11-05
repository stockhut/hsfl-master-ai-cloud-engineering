package reverse_proxy

import (
	"golang.org/x/exp/maps"
	"io"
	"log"
	"net/http"
	"strings"
)

type ReverseProxy struct {
	logger   *log.Logger
	services []Service
}

func (proxy *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handled := false
	for _, service := range proxy.services {
		if strings.HasPrefix(r.URL.Path, service.Route) {
			handled = true

			proxy.logger.Printf("%s => %s (%s)\n", r.URL, service.Name, service.TargetHost)
			err := forward(w, r, service)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				proxy.logger.Printf("Failed to forward request: %s", err)
			}
			break
		}
	}

	if !handled {
		proxy.logger.Printf("No matching service for %v\n", r.URL)
	}
}

func New(logger *log.Logger, services []Service) *ReverseProxy {
	return &ReverseProxy{
		logger:   logger,
		services: services,
	}
}

type Service struct {
	Name       string
	Route      string
	TargetHost string
}

// forward the request to the given service, writing the service response to w
func forward(w http.ResponseWriter, r *http.Request, service Service) error {

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
