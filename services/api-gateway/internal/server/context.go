package server

import (
	"context"
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/middleware"
)

type contextKey string

const authenticationContextKey = contextKey("authentication")

func contextSetAuthentication(r *http.Request, claims middleware.Claims) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), authenticationContextKey, claims))

}
func contextGetAuthentication(r *http.Request) middleware.Claims {
	claims, ok := r.Context().Value(authenticationContextKey).(middleware.Claims)
	if !ok {
		return middleware.Claims{}
	}

	return claims
}
