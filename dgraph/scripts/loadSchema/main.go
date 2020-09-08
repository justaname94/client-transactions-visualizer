package main

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// CancelFunc represents a function that should be deffered on its called after use
type CancelFunc func()

func connect() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return dgraphClient, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}
}

func main() {
	client, cancel := connect()
	defer cancel()

	loadSchema(client)
}

func loadSchema(client *dgo.Dgraph) {
	op := &api.Operation{
		Schema: `
		id:         string @index(exact)   .
		name:       string                 .
		age:        int                    .
		price:      int                    .
		buyerID:    string @index(exact)   .
		ip:         string @index(exact)   .
		device:     string                 .
		productIDs: [string] @index(exact) .
		date:       string @index(exact)   .
		transaction: [uid] @reverse        .
		product: [uid]     @reverse        .

		type Buyer {
			id:   string
			name: string 
			age:  int
		}
		
		type Product {
			id:    string
			name:  string
			price: int
		}
		
		type Transaction {
			id:         string
			buyerID:    string
			ip:         string
			device:     string
			productIDs: [string]
		}`,
	}

	ctx := context.Background()

	if err := client.Alter(ctx, op); err != nil {
		log.Fatalf("Error while mutating schema: %v\n", err)
	}
}
