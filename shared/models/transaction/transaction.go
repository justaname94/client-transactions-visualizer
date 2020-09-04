package models

import (
	"transactions/shared/utils"
)

// Transaction is the struct type for a transaction. A transaction is defined
// as a serie of product purchase at a moment by a Buyer.
type Transaction struct {
	ID         string   `json:"id,,omitempty"`
	BuyerID    string   `json:"buyerID,omitempty"`
	IP         string   `json:"ip,,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIDs []string `json:"productIDs,omitempty"`
}

// NewTransaction returns a transaction structure with the given values
func NewTransaction(id, buyerID, ip, device string,
	productIds []string) (*Transaction, error) {

	if id == "" {
		return nil, utils.ErrMissingField("id")
	}

	if buyerID == "" {
		return nil, utils.ErrMissingField("buyerId")
	}

	if ip == "" {
		return nil, utils.ErrMissingField("ip")
	}

	if device == "" {
		return nil, utils.ErrMissingField("device")
	}

	return &Transaction{
		ID:         id,
		BuyerID:    buyerID,
		IP:         ip,
		Device:     device,
		ProductIDs: productIds,
	}, nil
}
