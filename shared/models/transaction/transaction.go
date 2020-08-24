package models

import (
	"transactions/shared/utils"
)

// Transaction is the struct type for a transaction. A transaction is defined
// as a serie of product purchase at a moment by a Buyer.
type Transaction struct {
	ID         string   `json:"id"`
	BuyerID    string   `json:"buyerID"`
	IP         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIDs []string `json:"productIDs"`
}

// NewClient returns a client structure with the given values
func NewClient(id, buyerID, ip, device string,
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
