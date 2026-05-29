package domain

import "errors"

var (
	ErrInvalidEmail  = errors.New("invalid email")
	ErrInvalidUser   = errors.New("invalid user")
	ErrEmptyPassword = errors.New("password is empty")
	ErrUserNotFound  = errors.New("user not found")
)
