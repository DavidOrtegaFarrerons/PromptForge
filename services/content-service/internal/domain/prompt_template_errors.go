package domain

import "errors"

var (
	ErrContentEmpty = errors.New("the prompt has no content")
)
