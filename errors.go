package feature

import (
	"fmt"
)

type PropertyNotFoundErr struct {
	property string
}

func (e *PropertyNotFoundErr) Error() string {
	return fmt.Sprintf("'%s' property not found", e.property)
}

func PropertyNotFoundError(prop string) *PropertyNotFoundErr {
	return &PropertyNotFoundErr{
		property: prop,
	}
}
