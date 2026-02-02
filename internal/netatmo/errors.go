package netatmo

import (
	"errors"
	"fmt"
)

// Sentinel errors for common conditions
var (
	ErrNotConfigured    = errors.New("credentials not configured")
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrTokenExpired     = errors.New("token expired")
	ErrNoDevices        = errors.New("no devices found")
	ErrNetworkError     = errors.New("network error")
	ErrAPIError         = errors.New("API error")
)

// AuthError represents authentication-related errors with remediation hints
type AuthError struct {
	Message string
	Hint    string
	Err     error
}

func (e *AuthError) Error() string {
	return e.Message
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

// UserMessage returns a user-friendly message with hint
func (e *AuthError) UserMessage() string {
	if e.Hint != "" {
		return fmt.Sprintf("%s\n💡 Hint: %s", e.Message, e.Hint)
	}
	return e.Message
}

// NewAuthError creates a new authentication error
func NewAuthError(message, hint string, err error) *AuthError {
	return &AuthError{
		Message: message,
		Hint:    hint,
		Err:     err,
	}
}

// APIError represents errors from the Netatmo API
type APIError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// UserMessage returns a user-friendly message
func (e *APIError) UserMessage() string {
	switch e.StatusCode {
	case 401, 403:
		return fmt.Sprintf("Authentication failed (HTTP %d)\n💡 Hint: Run 'netatmo login' to re-authenticate", e.StatusCode)
	case 404:
		return "Resource not found. Check your Netatmo account has weather station devices."
	case 429:
		return "Rate limit exceeded. Please wait a moment and try again."
	case 500, 502, 503:
		return "Netatmo API is temporarily unavailable. Please try again later."
	default:
		return fmt.Sprintf("API request failed with status %d: %s", e.StatusCode, e.Message)
	}
}

// NewAPIError creates a new API error
func NewAPIError(statusCode int, message string, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// NetworkError represents network-related errors
type NetworkError struct {
	Message string
	Err     error
}

func (e *NetworkError) Error() string {
	return e.Message
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

// UserMessage returns a user-friendly message
func (e *NetworkError) UserMessage() string {
	return fmt.Sprintf("Network error: %s\n💡 Hint: Check your internet connection", e.Message)
}

// NewNetworkError creates a new network error
func NewNetworkError(message string, err error) *NetworkError {
	return &NetworkError{
		Message: message,
		Err:     err,
	}
}

// FormatError returns a user-friendly error message for known error types
func FormatError(err error) string {
	if err == nil {
		return ""
	}

	var authErr *AuthError
	if errors.As(err, &authErr) {
		return authErr.UserMessage()
	}

	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.UserMessage()
	}

	var netErr *NetworkError
	if errors.As(err, &netErr) {
		return netErr.UserMessage()
	}

	return err.Error()
}
