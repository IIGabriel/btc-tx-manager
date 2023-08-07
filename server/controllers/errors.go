package controllers

import (
	"fmt"
)

var (
	ErrRequired = func(element string) error { return fmt.Errorf("%s is required", element) }
)
