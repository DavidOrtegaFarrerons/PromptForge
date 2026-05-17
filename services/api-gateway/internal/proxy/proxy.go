package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
}

func New(target, prefix string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)

		if r.URL.Path == "" {
			r.URL.Path = "/"
		}

		log.Printf("Converted url path to: %s", r.URL.Path)

		reverseProxy.ServeHTTP(w, r)
	})

}
