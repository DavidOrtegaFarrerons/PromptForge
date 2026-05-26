package httpcontext

import (
	"context"
	"net/http"
)

type contextKey string

const userIDContextKey = contextKey("userID")

func ContextSetUserID(r *http.Request, userID string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userIDContextKey, userID))
}

func ContextGetUserID(r *http.Request) string {
	userID, ok := r.Context().Value(userIDContextKey).(string)
	if !ok {
		return ""
	}

	return userID
}
