// у этого класса должен быть клиенит
package mongo

import (
	"context"
	"log"

	"github.com/xxarchexx/tnvcache/internal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var store Store

//Store new Store
func Init(dbname, tableName, connectionstring string) (internal.CacheRepository, internal.IClose) {
	store = Store{}
	ctx := context.Background()
	options := options.Client().ApplyURI(connectionstring)
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		log.Fatal(err)
	}
	store.client = client
	store.collection = store.client.Database(dbname).Collection(tableName)

	var cacheRepo internal.CacheRepository
	cacheRepo = internal.CacheRepository(&store)

	var close internal.IClose
	close = internal.IClose(&store)

	return cacheRepo, close

}

//Close connection
func (store *Store) Close() {
	store.client.Disconnect(context.Background())
}

//Get data from mongo
func (store *Store) Get(key string) (string, error) {
	filter := bson.M{"key": key}

	type Result struct {
		Value string `bson:"value"`
	}

	var result Result
	store.collection.FindOne(context.Background(), filter).Decode(&result)
	return result.Value, nil

}

//Set data to mongo
func (store *Store) Set(key, value string) error {
	type saveStruct struct {
		Key   string
		Value string
	}
	newSaveStruct := saveStruct{Key: key, Value: value}

	_, err := store.collection.InsertOne(context.Background(), newSaveStruct)
	if err != nil {
		return err
	}

	return nil
}
