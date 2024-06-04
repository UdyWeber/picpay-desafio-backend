package errors

import (
	"fmt"
)

type BaseError struct {
	OriginalMessage string `json:"original_message,omitempty"`
	Message         string `json:"message"`
}

func (e *BaseError) Error() string {
	if e.OriginalMessage != "" {
		return fmt.Sprintf("[ERROR] %s. Caused by %s", e.Message, e.OriginalMessage)
	}

	return fmt.Sprintf("[ERROR] %s", e.Message)
}

func NewBaseError(originalMessage, message string) *BaseError {
	return &BaseError{
		OriginalMessage: originalMessage,
		Message:         message,
	}
}
