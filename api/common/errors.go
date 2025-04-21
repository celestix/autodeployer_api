package common

import "fmt"

type ErrorCode int

const (
	ErrGeneral ErrorCode = iota
	ErrUnauthorized
	ErrInvalidToken
	ErrExpiredToken
	ErrInvalidRequestBody
	ErrUserNotFound
	ErrInvalidPassword
	ErrUserAlreadyExists
)

var Errors = map[ErrorCode]string{
	ErrGeneral:            "Error occured",
	ErrUnauthorized:       "Unauthorized",
	ErrInvalidToken:       "Invalid token",
	ErrExpiredToken:       "Token expired",
	ErrInvalidRequestBody: "Invalid request body",
	ErrUserNotFound:       "User not found",
	ErrInvalidPassword:    "Invalid password",
	ErrUserAlreadyExists:  "The User already exists",
}

type ErrorEnum interface {
	ErrorCode | string
}

func ErrorMessage(message string, args ...any) Response {
	return Response{
		Error: &Error{
			Code:        ErrGeneral,
			Description: fmt.Sprintf(message, args...),
		},
	}
}

func ErrorResponse(code ErrorCode) Response {
	return Response{
		Error: &Error{
			Code:        code,
			Description: Errors[code],
		},
	}
}
