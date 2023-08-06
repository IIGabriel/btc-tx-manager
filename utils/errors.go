package utils

import (
	"errors"
)

var (
	ErrMissingEnv = errors.New("env not found or not loaded")
	ErrEnvType    = errors.New("this env not match with selected type")
)
