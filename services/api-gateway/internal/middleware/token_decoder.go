package middleware

import (
	"time"
)

type Claims struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}

type TokenDecoder interface {
	Decode(tokenString string) (Claims, error)
}
