package server

import (
	"net/http"
	"strings"
)

func (app *Application) DecodeTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}
		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			next.ServeHTTP(w, r)
			return
		}

		token := parts[1]

		claims, err := app.tokenDecoder.Decode(token)
		if err != nil {
			return
		}

		r = contextSetAuthentication(r, claims)
		r.Header.Set("X-User-ID", claims.UserID)
		r.Header.Set("X-User-Email", claims.Email)
		next.ServeHTTP(w, r)
	})
}
