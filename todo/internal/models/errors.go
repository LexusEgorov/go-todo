package models

import (
	"fmt"
)

func NewEmptyErr(field string) error {
	return fmt.Errorf("field '%s' is required", field)
}
