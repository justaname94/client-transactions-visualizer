package models

import (
	"errors"
	"transactions/shared/utils"
)

// Product is the struct type for a product
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// NewClient returns a client structure with the given values
func NewClient(id, name string, price int) (*Product, error) {
	if id == "" {
		return nil, utils.ErrMissingField("id")
	}

	if name == "" {
		return nil, utils.ErrMissingField("name")
	}

	if price <= 0 {
		return nil, errors.New("invalid number for price")
	}

	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}, nil
}
