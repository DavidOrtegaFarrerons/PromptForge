package routes

import (
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/config"
	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/proxy"
)

func Register(mux *http.ServeMux, cfg config.Config) {
	mux.Handle("/api/auth/", proxy.New(cfg.AuthServiceUrl, "/api/auth"))
	mux.Handle("/api/billing/", proxy.New(cfg.BillingServiceUrl, "/api/billing"))
	mux.Handle("/api/content/", proxy.New(cfg.ContentServiceUrl, "/api/content"))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
