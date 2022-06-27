package gobitkub

import (
	"errors"
	"fmt"
)

type BitkubError struct {
	Code int
	Err  error
}

func (e *BitkubError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func NewBitkubError(errorCode int) *BitkubError {
	return &BitkubError{
		Code: errorCode,
		Err:  errors.New(ERROR_MESSAGE[errorCode]),
	}
}
