package helper

import "github.com/mohammad-ebrahimi-it/car-shoping/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}
func GenerateBaseResponseWithAnyError(result any, success bool, resultCode int, err any) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.(error),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationError(err),
	}
}
