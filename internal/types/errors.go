package types

import (
	"net/http"
)

type StatusError struct {
	Message        string `json:"message"`
	Code           string `json:"code"`
	HTTPCode       int    `json:"-"`
	DisplayMessage string `json:"displayMessage"`
}

func (e *StatusError) Error() string {
	return e.Message
}

func NewValidationError(message string) *StatusError {
	return &StatusError{
		Message:  message,
		Code:     "error_processing_request",
		HTTPCode: http.StatusBadRequest,
	}
}

func NewInternalServerError() *StatusError {
	return &StatusError{
		Message:        "internal server error",
		HTTPCode:       http.StatusInternalServerError,
		DisplayMessage: "",
	}
}

func InternalServerError(err error) *StatusError {
	return &StatusError{
		Message:        err.Error(),
		HTTPCode:       http.StatusInternalServerError,
		DisplayMessage: "Something went wrong. Please try again",
	}
}

func NewBadRequestError() *StatusError {
	return &StatusError{
		Message:        "bad request",
		DisplayMessage: "",
		HTTPCode:       http.StatusBadRequest,
	}
}

func BadRequestError(err error) *StatusError {
	return &StatusError{
		Message:        err.Error(),
		DisplayMessage: "Something went wrong. Please try again",
		HTTPCode:       http.StatusBadRequest,
	}
}

func NewUnAuthorizedError() *StatusError {
	return &StatusError{
		Message:        "unauthorized",
		DisplayMessage: "",
		HTTPCode:       http.StatusUnauthorized,
	}
}

func NewNotFoundError() *StatusError {
	return &StatusError{
		Message:        "not found",
		DisplayMessage: "",
		HTTPCode:       http.StatusNotFound,
	}
}

func NotFoundError(err error) *StatusError {
	return &StatusError{
		Message:        err.Error(),
		DisplayMessage: "Something went wrong. Please try again",
		HTTPCode:       http.StatusNotFound,
	}
}

func BuildErrorResponse(err *StatusError) ErrorResponse {
	return ErrorResponse{
		Error: Error{
			DisplayMessage: err.DisplayMessage,
			Message:        err.Message,
			Code:           err.Code,
		},
	}
}
