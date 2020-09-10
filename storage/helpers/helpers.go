package helpers

import (
	"encoding/json"
	buyer "transactions/shared/models/buyer"
	product "transactions/shared/models/product"
	"transactions/shared/utils"
	"transactions/storage"

	"github.com/dgraph-io/dgo/v2"
)

// SaveEntities checks for duplicates, filters them and saves all the
// entities into the database
func SaveEntities(client *dgo.Dgraph, entities utils.Entities) error {

	// Filter duplicates
	filteredBuyers, err := filterBuyers(client, entities.Buyers)
	if err != nil {
		return err
	}

	filteredProducts, err := filterProducts(client, entities.Products)
	if err != nil {
		return err
	}

	buyersJSON, err := json.Marshal(filteredBuyers)
	if err != nil {
		return err
	}
	if err := storage.Save(client, buyersJSON); err != nil {
		return err
	}

	transactionsJSON, err := json.Marshal(entities.Transactions)
	if err != nil {
		return err
	}
	if err := storage.Save(client, transactionsJSON); err != nil {
		return err
	}

	productsJSON, err := json.Marshal(filteredProducts)
	if err != nil {
		return err
	}
	if err := storage.Save(client, productsJSON); err != nil {
		return err
	}

	return nil
}

// Check the current items on buyers and returns back the non duplicated ones
func filterBuyers(client *dgo.Dgraph,
	buyers []*buyer.Buyer) ([]*buyer.Buyer, error) {
	type Buyers struct {
		Buyers []buyer.Buyer `json:"buyers"`
	}

	buyersInDb, err := storage.Query(client, storage.AllBuyers,
		map[string]string{})
	if err != nil {
		return []*buyer.Buyer{}, err
	}

	var buyersArr Buyers
	var filteredBuyers []*buyer.Buyer
	buyersMap := make(map[string]bool)
	if err = json.Unmarshal(buyersInDb.Json, &buyersArr); err != nil {
		return []*buyer.Buyer{}, err
	}

	for _, buyer := range buyersArr.Buyers {
		buyersMap[buyer.ID] = true
	}

	for _, buyer := range buyers {
		if _, exist := buyersMap[buyer.ID]; exist {
			continue
		}
		filteredBuyers = append(filteredBuyers, buyer)
	}

	return filteredBuyers, nil
}

// Check the current items on products and returns back the non duplicated ones.
// Until now works like filterBuyers but that might change in  the future.
func filterProducts(client *dgo.Dgraph,
	products []*product.Product) ([]*product.Product, error) {
	type Products struct {
		Products []product.Product `json:"products"`
	}

	productsInDb, err := storage.Query(client, storage.AllProducts,
		map[string]string{})
	if err != nil {
		return []*product.Product{}, err
	}

	var productsArr Products
	var filteredProducts []*product.Product
	productsMap := make(map[string]bool)

	if err = json.Unmarshal(productsInDb.Json, &productsArr); err != nil {
		return []*product.Product{}, err
	}

	for _, product := range productsArr.Products {
		productsMap[product.ID] = true
	}

	for _, product := range products {
		if _, exist := productsMap[product.ID]; exist {
			continue
		}
		filteredProducts = append(filteredProducts, product)
	}

	return filteredProducts, nil
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
