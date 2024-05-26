package domain

import (
	"strconv"
	"strings"
)

const (
	ErrUnauthorizedCode   Code = 401
	ErrNotFoundCode       Code = 404
	ErrRequestTimeoutCode Code = 408
	ErrConflictCode       Code = 409

	ErrInternalCode Code = 500

	ErrUnauthorizedMessage = "Unauthorized"
	ErrInternalMessage     = "Internal error"
	ErrConflictMessage     = "Entity conflict"
	ErrNotFoundMessage     = "Not found"
)

type Code int

type Error struct {
	Code    Code
	Message string
}

func (e *Error) Error() string {
	var err strings.Builder
	err.WriteString("code: ")
	err.WriteString(strconv.Itoa(int(e.Code)))
	err.WriteString("message: ")
	err.WriteString(e.Message)

	return err.String()
}

func NewErr(code Code, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
