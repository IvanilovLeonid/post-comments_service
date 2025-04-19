package apperrors

import (
	"encoding/json"
	"fmt"
)

// APIError представляет структурированную ошибку API
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error реализует интерфейс error
func (e APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// WithDetails добавляет детали к ошибке
func (e APIError) WithDetails(details string) APIError {
	e.Details = details
	return e
}

// MarshalJSON кастомная маршализация ошибки
func (e APIError) MarshalJSON() ([]byte, error) {
	type Alias APIError
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(e),
	})
}

// Predefined errors
var (
	ErrBadRequest   = APIError{Code: "BAD_REQUEST", Message: "Invalid request"}
	ErrUnauthorized = APIError{Code: "UNAUTHORIZED", Message: "Authentication required"}
	ErrNotFound     = APIError{Code: "NOT_FOUND", Message: "Resource not found"}
	// ... другие предопределенные ошибки
)

// New создает новую ошибку API
func New(code, message string) APIError {
	return APIError{
		Code:    code,
		Message: message,
	}
}

func (e *APIError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
