package messages

import (
	"errors"
)

// ErrMissingField servers as a wrapper to provide an error on missing model
// fields
func ErrMissingField(field string) error {
	return errors.New("missing parameter " + field)
}

// ErrNotFound servers as a wrapper to provide an error on not found
//entities
func ErrNotFound(entity string) error {
	return errors.New(entity + " not found")
}
