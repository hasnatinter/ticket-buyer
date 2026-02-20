package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors []string `json:""errors`
}

func ToErrResponse(err error) *ErrorResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrorResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be maxmimum of length %s", err.Field(), err.Param())
			case "min":
				resp.Errors[i] = fmt.Sprintf("%s must be minumum of length %s", err.Field(), err.Param())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("%s must be a valid date", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("%s must follow format %s", err.Field(), err.Param())
				}
			case "required_with":
				resp.Errors[i] = fmt.Sprintf("%s is a required with %s", err.Field(), err.Param())
			case "alphanum":
				resp.Errors[i] = fmt.Sprintf("%s must be alphanumeric", err.Field())
			}
		}
		return &resp
	}
	return nil
}
