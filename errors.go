package feature

import (
	"fmt"
)

type NotImplementedError struct {
}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("Not implemented")
}

func (e *NotImplementedError) String() string {
	return e.Error()
}

type NotFoundError struct {
	path string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Property '%s' not found", e.path)
}

func (e *NotFoundError) String() string {
	return e.Error()
}
