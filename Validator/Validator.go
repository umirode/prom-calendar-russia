package Validator

import (
	"github.com/umirode/prom-calendar-russia/Validator/Custom"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		Validator: validator.New(),
	}

	v.registerCustomValidator(Custom.NewDateCustomValidator())

	return v
}

func (v *Validator) registerCustomValidator(customValidator Custom.ICustomValidator) {
	v.Validator.RegisterValidation(customValidator.GetTag(), customValidator.GetValidator)
}
