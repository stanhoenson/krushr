package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var globalValidator *validator.Validate

func customValidation(fl validator.FieldLevel) bool {
	// custom validation logic
	return true // or false
}

func InitializeValidators() {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		globalValidator = validator.New()
		globalValidator.RegisterValidation("custom", customValidation)
	}
}

func ValidateStruct(structToValidate interface{}) error {
	return globalValidator.Struct(structToValidate)
}

func ValidateField(value interface{}, tag string) error {
	return globalValidator.Var(value, tag)
}
