package mongo

import (
	"context"
	"log"

	"github.com/xxarchexx/tnvcache/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectionRepositry
type ConnectionRepositry MongoStore

//GetConnectionRepository  for db
func (store *ConnectionRepositry) GetConnectionRepository() internal.ConnectionRepository {
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

	store.Client = client
	collection = client.Database(dbname).Collection(tableName)
	return nil
}

//SetConnection set exists connection
func (store *ConnectionRepositry) SetConnection(dbname, tableName string, client interface{}) {
	store.Client = client.(*mongo.Client)
	collection = store.Client.Database(dbname).Collection(tableName)
}

//CloseConnection close connection
func (store *ConnectionRepositry) CloseConnection() error {
	ctx := context.Background()
	err := store.Client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
