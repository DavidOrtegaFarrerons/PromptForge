package domain

import "errors"

var (
	ErrPromptIDEmpty  = errors.New("prompt id must be provided")
	ErrOwnerIdEmpty   = errors.New("owner id must be provided")
	ErrTitleMinLength = errors.New("title must be at least 4 characters long")
	ErrTitleMaxLength = errors.New("title must be at maximum 64 characters long")
)
