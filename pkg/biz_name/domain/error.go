package domain

import (
	"fmt"
	"net/http"
)

// AppError handles application exception.
type AppError struct {
	HTTPStatus int
	Label      string
	Message    string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s - %s", e.Label, e.Message)
}

// NewAppError New functions create a new AppError instance
func NewAppError(status int, label, message string) *AppError {
	return &AppError{
		HTTPStatus: status,
		Label:      label,
		Message:    message,
	}
}

var (
	ErrInternal = &AppError{HTTPStatus: http.StatusInternalServerError, Label: "INTERNAL", Message: "Internal error"}
	ErrNotFound = &AppError{HTTPStatus: http.StatusNotFound, Label: "Not Found", Message: "Resource not found"}
)
