package models

import (
	"errors"
	"transactions/shared/utils"
)

// Buyer is the struct type for a buyer
type Buyer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NewBuyer returns a client structure with the given values
func NewBuyer(id, name string, age int) (*Buyer, error) {
	if id == "" {
		return nil, utils.ErrMissingField("id")
	}

	if name == "" {
		return nil, utils.ErrMissingField("name")
	}

	if age <= 0 {
		return nil, errors.New("invalid number for age")
	}

	return &Buyer{
		ID:   id,
		Name: name,
		Age:  age,
	}, nil
}
