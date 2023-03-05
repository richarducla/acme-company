package errors

import (
	"fmt"
	"net/http"
)

var (
	BadRequestErr          = (&HandlerError{}).WithMessage("bad-request-error").WithStatusCode(http.StatusBadRequest)
	UnauthorizedErr        = (&HandlerError{}).WithMessage("unauthorized-error").WithStatusCode(http.StatusUnauthorized)
	UnprocessableEntityErr = (&HandlerError{}).WithMessage("unprocessable-entity-error").WithStatusCode(http.StatusUnprocessableEntity)
	InternalServerErrorErr = (&HandlerError{}).WithMessage("internal-server-error").WithStatusCode(http.StatusInternalServerError)
)

// HandlerError is the error response format
type HandlerError struct {
	StatusCode  int    `json:"statusCode"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (e HandlerError) Error() string {
	if e.Message == "" {
		return ""
	}

	return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Message, e.Description)
}

func (e *HandlerError) WithStatusCode(statusCode int) *HandlerError {
	e.StatusCode = statusCode
	return e
}

func (e *HandlerError) WithDescription(description string) *HandlerError {
	e.Description = description
	return e
}

func (e *HandlerError) WithMessage(message string) *HandlerError {
	e.Message = message
	return e
}
