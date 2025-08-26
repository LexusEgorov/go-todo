package models

import (
	"errors"
)

const (
	ErrGetBody  = "Error while reading body"
	ErrReadJSON = "JSON is invalid"
)

var (
	ErrBadBody  = errors.New("body isn't valid")
	ErrNotFound = errors.New("not found")
)
