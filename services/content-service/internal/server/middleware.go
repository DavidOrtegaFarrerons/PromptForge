package server

import (
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/httpcontext"
)

func (app *Application) GetGatewayHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		if userID != "" {
			r = httpcontext.ContextSetUserID(r, userID)
		}

		next.ServeHTTP(w, r)
	})
}
