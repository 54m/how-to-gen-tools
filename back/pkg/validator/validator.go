package validator

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

// Validator - validator
type Validator struct {
	validator *validator.Validate
}

// NewValidator - constructor
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate - validate
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		return xerrors.Errorf("validation error: %w", err)
	}

	return nil
}

// DefaultValidator - default validator
var DefaultValidator = NewValidator()

// Validate - validate
func Validate(i interface{}) error {
	return DefaultValidator.Validate(i)
}
