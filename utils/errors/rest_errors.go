package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode"`
	Error   string `json:"error"`
}

func BadRequestError(errorMessage string) *RestError {
	return &RestError{
		Message: errorMessage,
		Code:    http.StatusBadRequest,
		Error:   "invalid json body",
	}
}

func NotFoundError(errorMessage string) *RestError {
	return &RestError{
		Message: errorMessage,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}
