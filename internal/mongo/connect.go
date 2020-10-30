package mongo

import (
	"context"
	"log"

	"github.com/xxarchexx/tnvcache/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectionRepositry
type ConnectionRepositry Store

//GetConnectionRepository  for db
func (store *Store) GetConnectionRepository() internal.ConnectionRepository {
	return (*ConnectionRepositry)(store)
}

//NewConnection new  connect to mongo
func (store *ConnectionRepositry) NewConnection(dbname, tableName, uri string) error {
	ctx := context.Background()
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal(err)
		return err
	}

	store.client = client
	collection = client.Database(dbname).Collection(tableName)
	return nil
}

//SetConnection set exists connection
func (store *ConnectionRepositry) SetConnection(dbname, tableName string, client interface{}) {
	store.client = client.(*mongo.Client)
	collection = store.client.Database(dbname).Collection(tableName)
}

//CloseConnection close connection
func (store *ConnectionRepositry) CloseConnection() error {
	ctx := context.Background()
	err := store.client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
