package main

import "github.com/go-playground/validator"

// Validator - validator struct
type Validator struct {
	validator *validator.Validate
}

// Validate - validates the specified input
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
