package commons

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func Validate(data interface{}) error {
	validate = validator.New()
	errs := validate.Struct(data)
	if errs == nil {
		return nil
	}
	errTip := errorData(errs.(validator.ValidationErrors))
	return errors.New(errTip)
}

func errorData(errs []validator.FieldError) string {
	for _, err := range errs {
		return fmt.Sprintf("%s is %s %s", err.Field(), err.Tag(), err.Param())
	}
	return "unknown error"
}
