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

func (e *APIError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
