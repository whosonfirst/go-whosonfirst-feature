package properties

import (
	"fmt"
)

type NotFoundError struct {
	path string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Propery '%s' not found", e.path)
}

func (e *NotFoundError) String() string {
	return e.Error()
}
