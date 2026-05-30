package domain

import "errors"

var ErrPromptLimitReached = errors.New("you have reached the maximum amount of prompts for your current plan")
