package models

import "fmt"

const (
	ErrGetBody  = "Error while reading body"
	ErrReadJSON = "JSON is invalid"
)

var (
	ErrBadBody  = fmt.Errorf("body isn't valid")
	ErrNotFound = fmt.Errorf("not found")
)
