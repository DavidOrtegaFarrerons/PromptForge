package application

import "errors"

var ErrDuplicateEmail = errors.New("the email is already in use")
