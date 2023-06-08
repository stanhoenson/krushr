package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func customValidation(fl validator.FieldLevel) bool {
	// custom validation logic
	return true // or false
}

func InitializeValidators() {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v := validator.New()
		v.RegisterValidation("custom", customValidation)
	}
}
