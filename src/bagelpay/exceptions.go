package bagelpay

import (
	"fmt"
	"net/http"
)

// BagelPayError represents a base error type for all BagelPay SDK errors
type BagelPayError struct {
	Message string
	Cause   error
}

func (e *BagelPayError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("BagelPay error: %s (caused by: %v)", e.Message, e.Cause)
	}
	return fmt.Sprintf("BagelPay error: %s", e.Message)
}

func (e *BagelPayError) Unwrap() error {
	return e.Cause
}

// NewBagelPayError creates a new BagelPayError
func NewBagelPayError(message string, cause error) *BagelPayError {
	return &BagelPayError{
		Message: message,
		Cause:   cause,
	}
}

// BagelPayAPIError represents an API-specific error
type BagelPayAPIError struct {
	*BagelPayError
	StatusCode int
	ErrorCode  string
	APIError   *APIError
}

func (e *BagelPayAPIError) Error() string {
	if e.APIError != nil {
		return fmt.Sprintf("BagelPay API error (status %d): %s", e.StatusCode, e.APIError.Message)
	}
	return fmt.Sprintf("BagelPay API error (status %d): %s", e.StatusCode, e.Message)
}

// String returns a formatted string representation of the error (equivalent to TypeScript toString)
func (e *BagelPayAPIError) String() string {
	parts := []string{e.Message}
	if e.StatusCode > 0 {
		parts = append(parts, fmt.Sprintf("Status: %d", e.StatusCode))
	}
	if e.ErrorCode != "" {
		parts = append(parts, fmt.Sprintf("Code: %s", e.ErrorCode))
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += " | " + parts[i]
	}
	return result
}

// NewBagelPayAPIError creates a new BagelPayAPIError
func NewBagelPayAPIError(statusCode int, apiError *APIError, cause error) *BagelPayAPIError {
	message := "API request failed"
	errorCode := ""
	if apiError != nil {
		message = apiError.Message
		if apiError.Code != 0 {
			errorCode = fmt.Sprintf("%d", apiError.Code)
		}
	}

	return &BagelPayAPIError{
		BagelPayError: NewBagelPayError(message, cause),
		StatusCode:    statusCode,
		ErrorCode:     errorCode,
		APIError:      apiError,
	}
}

// BagelPayAuthenticationError represents authentication-related errors
type BagelPayAuthenticationError struct {
	*BagelPayAPIError
}

func (e *BagelPayAuthenticationError) Error() string {
	return fmt.Sprintf("BagelPay authentication error: %s", e.Message)
}

// NewBagelPayAuthenticationError creates a new BagelPayAuthenticationError
func NewBagelPayAuthenticationError(message string, statusCode int, errorCode string, apiError *APIError, cause error) *BagelPayAuthenticationError {
	if statusCode == 0 {
		statusCode = http.StatusUnauthorized
	}
	return &BagelPayAuthenticationError{
		BagelPayAPIError: &BagelPayAPIError{
			BagelPayError: NewBagelPayError(message, cause),
			StatusCode:    statusCode,
			ErrorCode:     errorCode,
			APIError:      apiError,
		},
	}
}

// NewBagelPayAuthenticationErrorSimple creates a new BagelPayAuthenticationError with minimal parameters
func NewBagelPayAuthenticationErrorSimple(message string, cause error) *BagelPayAuthenticationError {
	return NewBagelPayAuthenticationError(message, http.StatusUnauthorized, "", nil, cause)
}

// BagelPayValidationError represents validation-related errors
type BagelPayValidationError struct {
	*BagelPayAPIError
}

func (e *BagelPayValidationError) Error() string {
	return fmt.Sprintf("BagelPay validation error: %s", e.Message)
}

// NewBagelPayValidationError creates a new BagelPayValidationError
func NewBagelPayValidationError(message string, statusCode int, errorCode string, apiError *APIError, cause error) *BagelPayValidationError {
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}
	return &BagelPayValidationError{
		BagelPayAPIError: &BagelPayAPIError{
			BagelPayError: NewBagelPayError(message, cause),
			StatusCode:    statusCode,
			ErrorCode:     errorCode,
			APIError:      apiError,
		},
	}
}

// NewBagelPayValidationErrorSimple creates a new BagelPayValidationError with minimal parameters
func NewBagelPayValidationErrorSimple(message string, cause error) *BagelPayValidationError {
	return NewBagelPayValidationError(message, http.StatusBadRequest, "", nil, cause)
}

// BagelPayNotFoundError represents not found errors
type BagelPayNotFoundError struct {
	*BagelPayAPIError
}

func (e *BagelPayNotFoundError) Error() string {
	return fmt.Sprintf("BagelPay not found error: %s", e.Message)
}

// NewBagelPayNotFoundError creates a new BagelPayNotFoundError
func NewBagelPayNotFoundError(message string, statusCode int, errorCode string, apiError *APIError, cause error) *BagelPayNotFoundError {
	if statusCode == 0 {
		statusCode = http.StatusNotFound
	}
	return &BagelPayNotFoundError{
		BagelPayAPIError: &BagelPayAPIError{
			BagelPayError: NewBagelPayError(message, cause),
			StatusCode:    statusCode,
			ErrorCode:     errorCode,
			APIError:      apiError,
		},
	}
}

// NewBagelPayNotFoundErrorSimple creates a new BagelPayNotFoundError with minimal parameters
func NewBagelPayNotFoundErrorSimple(message string, cause error) *BagelPayNotFoundError {
	return NewBagelPayNotFoundError(message, http.StatusNotFound, "", nil, cause)
}

// BagelPayRateLimitError represents rate limit errors
type BagelPayRateLimitError struct {
	*BagelPayAPIError
}

func (e *BagelPayRateLimitError) Error() string {
	return fmt.Sprintf("BagelPay rate limit error: %s", e.Message)
}

// NewBagelPayRateLimitError creates a new BagelPayRateLimitError
func NewBagelPayRateLimitError(message string, statusCode int, errorCode string, apiError *APIError, cause error) *BagelPayRateLimitError {
	if statusCode == 0 {
		statusCode = http.StatusTooManyRequests
	}
	return &BagelPayRateLimitError{
		BagelPayAPIError: &BagelPayAPIError{
			BagelPayError: NewBagelPayError(message, cause),
			StatusCode:    statusCode,
			ErrorCode:     errorCode,
			APIError:      apiError,
		},
	}
}

// NewBagelPayRateLimitErrorSimple creates a new BagelPayRateLimitError with minimal parameters
func NewBagelPayRateLimitErrorSimple(message string, cause error) *BagelPayRateLimitError {
	return NewBagelPayRateLimitError(message, http.StatusTooManyRequests, "", nil, cause)
}

// BagelPayServerError represents server-side errors
type BagelPayServerError struct {
	*BagelPayAPIError
}

func (e *BagelPayServerError) Error() string {
	return fmt.Sprintf("BagelPay server error: %s", e.Message)
}

// NewBagelPayServerError creates a new BagelPayServerError
func NewBagelPayServerError(message string, statusCode int, errorCode string, apiError *APIError, cause error) *BagelPayServerError {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return &BagelPayServerError{
		BagelPayAPIError: &BagelPayAPIError{
			BagelPayError: NewBagelPayError(message, cause),
			StatusCode:    statusCode,
			ErrorCode:     errorCode,
			APIError:      apiError,
		},
	}
}

// NewBagelPayServerErrorSimple creates a new BagelPayServerError with minimal parameters
func NewBagelPayServerErrorSimple(statusCode int, message string, cause error) *BagelPayServerError {
	return NewBagelPayServerError(message, statusCode, "", nil, cause)
}

// IsAuthenticationError checks if the error is an authentication error
func IsAuthenticationError(err error) bool {
	_, ok := err.(*BagelPayAuthenticationError)
	return ok
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*BagelPayValidationError)
	return ok
}

// IsNotFoundError checks if the error is a not found error
func IsNotFoundError(err error) bool {
	_, ok := err.(*BagelPayNotFoundError)
	return ok
}

// IsRateLimitError checks if the error is a rate limit error
func IsRateLimitError(err error) bool {
	_, ok := err.(*BagelPayRateLimitError)
	return ok
}

// IsServerError checks if the error is a server error
func IsServerError(err error) bool {
	_, ok := err.(*BagelPayServerError)
	return ok
}

// IsAPIError checks if the error is any API error
func IsAPIError(err error) bool {
	_, ok := err.(*BagelPayAPIError)
	return ok
}
