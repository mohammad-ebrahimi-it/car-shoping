package helper

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	service_errors.OtpExists:  409,
	service_errors.OtpUsed:    409,
	service_errors.OtpInValid: 400,

	service_errors.EmailExists:      409,
	service_errors.UsernameExists:   409,
	service_errors.RecordNotFound:   404,
	service_errors.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]

	if !ok {
		return http.StatusInternalServerError
	}

	return value
}
