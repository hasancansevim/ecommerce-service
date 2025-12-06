package _errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

func NewBadRequest(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFound(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(err error) *AppError {
	return &AppError{
		Code:     http.StatusInternalServerError,
		Message:  "Sunucu Hatası",
		Internal: err,
	}
}
