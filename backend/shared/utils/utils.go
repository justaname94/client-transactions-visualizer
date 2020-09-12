package utils

import (
	buyer "transactions/shared/models/buyer"
	product "transactions/shared/models/product"
	transaction "transactions/shared/models/transaction"
)

// Entities represents an array of the used models
type Entities struct {
	Buyers       []*buyer.Buyer
	Products     []*product.Product
	Transactions []*transaction.Transaction
}
