package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode"`
	Error   string `json:"error"`
}

func BadRequestError(errorMessage string) *RestError {
	return &RestError{
		Message: "invalid json body",
		Code:    http.StatusBadRequest,
		Error:   errorMessage,
	}
}
