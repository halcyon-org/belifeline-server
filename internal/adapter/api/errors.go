package api

import (
	"fmt"

	"github.com/halcyon-org/kizuna/internal/domain/domain"
)

type ValidationError struct {
	Message string
}

func NewValidationError(message string) error {
	return &ValidationError{Message: message}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s %s", e.Message, domain.Face)
}
