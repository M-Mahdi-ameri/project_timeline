package utils

import "github.com/go-playground/validator"

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}
