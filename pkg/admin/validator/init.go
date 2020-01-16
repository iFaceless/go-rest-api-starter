package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()
}
