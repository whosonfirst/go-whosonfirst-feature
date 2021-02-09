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
