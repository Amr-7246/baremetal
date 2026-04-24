package domain

import "errors"

var ( //? Teach me the var() syntax 
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrWeakPassword      = errors.New("password too weak")
	ErrInvalidUserID     = errors.New("invalid user ID")
)
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{Field: field, Message: message}
}