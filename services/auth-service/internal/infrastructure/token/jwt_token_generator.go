package token

import (
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenGenerator struct {
	signingSecret  []byte
	expirationTime time.Duration
}

func NewJwtTokenGenerator(signingSecret []byte, expirationTime time.Duration) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		signingSecret:  signingSecret,
		expirationTime: expirationTime,
	}
}

func (t *JwtTokenGenerator) Generate(userID domain.UserID, email domain.Email) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID,
		"email": email.Value(),
		"exp":   time.Now().Add(t.expirationTime).Unix(),
	}).SignedString(t.signingSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}
