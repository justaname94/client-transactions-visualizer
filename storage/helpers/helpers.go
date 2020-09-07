package helpers

import (
	"encoding/json"
	"transactions/shared/utils"
	"transactions/storage"

	"github.com/dgraph-io/dgo/v2"
)

// SaveEntities saves all the entities into the database
func SaveEntities(client *dgo.Dgraph, entities utils.Entities) error {
	buyersJSON, err := json.Marshal(entities.Buyers)
	if err != nil {
		return err
	}
	if err := storage.Save(client, buyersJSON); err != nil {
		return err
	}

	productsJSON, err := json.Marshal(entities.Products)
	if err != nil {
		return err
	}
	if err := storage.Save(client, productsJSON); err != nil {
		return err
	}

	transactionsJSON, err := json.Marshal(entities.Transactions)
	if err != nil {
		return err
	}
	if err := storage.Save(client, transactionsJSON); err != nil {
		return err
	}

	return nil
}

// ConnectFields bulk connects all the transactions with their buyers
// and the products with their transactions
func ConnectFields(client *dgo.Dgraph, entities utils.Entities) error {
	var buyerArr []string

	for _, buyer := range entities.Buyers {
		buyerArr = append(buyerArr, buyer.ID)
	}

	if err := storage.BulkConnect(client, "id", "buyerID", "transaction",
		buyerArr); err != nil {
		return err
	}

	var productArr []string

	for _, product := range entities.Products {
		productArr = append(productArr, product.ID)
	}

	if err := storage.BulkConnect(client, "productIDs", "id", "product",
		productArr); err != nil {
		return err
	}

	return nil
}
