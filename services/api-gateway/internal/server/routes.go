package server

import (
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/proxy"
)

func (app *Application) routes(mux *http.ServeMux) {
	mux.Handle("/api/auth/", app.DecodeTokenMiddleware(proxy.New(app.cfg.AuthServiceUrl, "/api/auth")))
	mux.Handle("/api/billing/", app.DecodeTokenMiddleware(proxy.New(app.cfg.BillingServiceUrl, "/api/billing")))
	mux.Handle("/api/content/", app.DecodeTokenMiddleware(proxy.New(app.cfg.ContentServiceUrl, "/api/content")))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
