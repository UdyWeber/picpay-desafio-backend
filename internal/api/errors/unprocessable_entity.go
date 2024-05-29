package errors

type UnprocessableEntityError struct {
	BaseError
	Fields map[string]string `json:"fields"`
}

func NewUnprocessableEntityError(errorMessage string, message string, fields map[string]string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		BaseError: BaseError{
			OriginalMessage: errorMessage,
			Message:         message,
		},
		Fields: fields,
	}
}

func (e *UnprocessableEntityError) ToResponse() []byte {
	return marshalApiError(e, "UnprocessableEntityError")
}
