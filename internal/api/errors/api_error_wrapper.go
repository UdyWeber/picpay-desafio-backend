package errors

type APIErrorWrapper struct {
	Code         int
	WrappedError error
}

func NewErrorWrapper(code int, err error) *APIErrorWrapper {
	return &APIErrorWrapper{code, err}
}
