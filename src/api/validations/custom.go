package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Value    string `json:"value"`
	Tag      string `json:"tag"`
	Message  string `json:"message"`
}

func GetValidationError(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, v := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = v.Field()
			el.Tag = v.Tag()
			el.Value = v.Param()
			validationErrors = append(validationErrors, el)
		}

		return &validationErrors
	}

	return nil
}
