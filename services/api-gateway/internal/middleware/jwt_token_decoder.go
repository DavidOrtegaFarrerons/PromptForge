package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenDecoder struct {
	secret string
}

func NewJwtTokenDecoder(secret string) *JwtTokenDecoder {
	return &JwtTokenDecoder{secret: secret}
}

func (d *JwtTokenDecoder) Decode(tokenString string) (Claims, error) {
	var jwtClaims struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (any, error) {
		return []byte(d.secret), nil
	})

	if err != nil {
		return Claims{}, err
	} else if !token.Valid {
		return Claims{}, ErrTokenNotValid
	}

	return Claims{
		UserID:    jwtClaims.Subject,
		Email:     jwtClaims.Email,
		ExpiresAt: jwtClaims.ExpiresAt.Time,
	}, nil
}
