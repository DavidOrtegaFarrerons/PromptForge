package application

import "errors"

var ErrDuplicateEmail = errors.New("the email is already in use")
var ErrTokenCouldNotBeGenerated = errors.New("token could not be generated")
