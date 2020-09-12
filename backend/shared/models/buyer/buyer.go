package models

import (
	"errors"
	messages "transactions/shared/error-messages"
)

// Buyer is the struct type for a buyer
type Buyer struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   int      `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

// NewBuyer returns a client structure with the given values
func NewBuyer(id, name string, age int) (*Buyer, error) {
	if id == "" {
		return nil, messages.ErrMissingField("id")
	}

	if name == "" {
		return nil, messages.ErrMissingField("name")
	}

	if age <= 0 {
		return nil, errors.New("invalid number for age")
	}

	return &Buyer{
		ID:    id,
		Name:  name,
		Age:   age,
		DType: []string{"Buyer"},
	}, nil
}
