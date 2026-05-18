package domain

import "errors"

var ErrInvalidEmail = errors.New("invalid email")
var ErrInvalidUser = errors.New("invalid user")
var ErrEmptyPassword = errors.New("password is empty")
var ErrTokenCouldNotBeGenerated = errors.New("token could not be generated")
var ErrUserNotFound = errors.New("user not found")
