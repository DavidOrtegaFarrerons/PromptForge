package domain

import (
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(strings.ToLower(value))

	if value == "" {
		return Email{}, ErrInvalidEmail
	}

	_, err := mail.ParseAddress(value)
	if err != nil {
		return Email{}, ErrInvalidEmail
	}

	return Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) IsZero() bool {
	return e.value == ""
}
