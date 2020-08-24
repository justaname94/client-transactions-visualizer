package utils

import (
	"errors"
)

// ErrMissingField servers as a wrapper to provide an error on missing model
// fields
func ErrMissingField(field string) error {
	return errors.New("missing parameter " + field)
}
