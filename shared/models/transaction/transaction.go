package models

import (
	"log"
	"strings"
	messages "transactions/shared/error-messages"
)

// Transaction is the struct type for a transaction. A transaction is defined
// as a serie of product purchase at a moment by a Buyer.
type Transaction struct {
	ID         string   `json:"id,,omitempty"`
	BuyerID    string   `json:"buyerID,omitempty"`
	IP         string   `json:"ip,,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIDs []string `json:"productIDs,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

// NewTransaction returns a transaction structure with the given values
func NewTransaction(id, buyerID, ip, device string,
	productIds []string) (*Transaction, error) {

	if id == "" {
		return nil, messages.ErrMissingField("id")
	}

	if buyerID == "" {
		return nil, messages.ErrMissingField("buyerId")
	}

	if ip == "" {
		return nil, messages.ErrMissingField("ip")
	}

	if device == "" {
		return nil, messages.ErrMissingField("device")
	}

	return &Transaction{
		ID:         id,
		BuyerID:    buyerID,
		IP:         ip,
		Device:     device,
		ProductIDs: productIds,
		DType:      []string{"Transaction"},
	}, nil
}

// ParseTransactions splits # as new line, '\x00' for inside characters
func ParseTransactions(data string) ([]*Transaction, error) {
	items := strings.Split(data, "#")

	var transactions []*Transaction

	for idx, item := range items {
		// Empty item
		if idx == 0 {
			continue
		}
		fields := strings.Split(item, "\x00")

		if len(fields) < 4 {
			// Log and ignore transaction error
			log.Printf("invalid transaction: %s\n", fields)
			continue
		}

		id := fields[0]
		buyerID := fields[1]
		ip := fields[2]
		device := fields[3]
		// Remove parentheses before splitting
		productIds := strings.Split(fields[4][1:len(fields[4])-1], ",")

		transaction, err := NewTransaction(id,
			buyerID, ip, device, productIds)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
