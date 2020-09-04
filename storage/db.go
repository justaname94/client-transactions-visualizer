package storage

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// CancelFunc represents a function that should be deffered on its called after
// use
type CancelFunc func()

// Connect returns a new connection to a local graphdb database
func Connect() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial(os.Getenv("DB_URL"), grpc.WithInsecure())

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
