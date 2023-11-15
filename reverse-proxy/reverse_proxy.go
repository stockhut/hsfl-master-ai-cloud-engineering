package reverse_proxy

import (
	"golang.org/x/exp/maps"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ReverseProxy struct {
	logger   *log.Logger
	services []Service
}

func (proxy *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	service := pickService(proxy.services, r.URL)
	if service == nil {
		proxy.logger.Printf("No matching service for %v\n", r.URL)
		return
	}

	proxy.logger.Printf("%s => %s (%s)\n", r.URL, service.Name, service.TargetHost)
	err := Forward(w, r, service.TargetHost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		proxy.logger.Printf("Failed to forward request: %s", err)
	}
}

func pickService(services []Service, u *url.URL) *Service {
	for _, service := range services {
		if strings.HasPrefix(u.Path, service.Route) {
			return &service
		}
	}

	return nil
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

// Forward the request to the given service, writing the service response to w
func Forward(w http.ResponseWriter, r *http.Request, host string) error {

	r.Host = host

	newUrl := "http://" + host + r.URL.Path

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
