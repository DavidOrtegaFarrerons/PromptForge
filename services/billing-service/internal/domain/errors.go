package domain

import "errors"

var (
	ErrEmptyAccount         = errors.New("account cannot be empty")
	ErrEmptyUserID          = errors.New("user id cannot be empty")
	ErrPlanInvalid          = errors.New("plan must be either free or pro")
	ErrAccountAlreadyExists = errors.New("account already exists")
)
