package controllers

import (
	"fmt"
)

var (
	ErrInvalid  = func(element string) error { return fmt.Errorf("invalid %s", element) }
	ErrRequired = func(element string) error { return fmt.Errorf("%s is required", element) }
)
