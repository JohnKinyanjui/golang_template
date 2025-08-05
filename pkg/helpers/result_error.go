package helpers

type ResultError struct {
	Understandable string
	Error          error
}

func Error(understandable string, err error) *ResultError {
	return &ResultError{
		Understandable: understandable,
		Error:          err,
	}
}
