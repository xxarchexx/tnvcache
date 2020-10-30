// у этого класса должен быть клиенит
package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const keyFieldName string = "key"
const valueParam string = "value"

var collection *mongo.Collection

type Store struct {
	client *mongo.Client
}

//Store new Store
func Init() *Store {
	return &Store{}
}

//Get data from mongo
func (store *Store) Get(key string) (string, error) {
	filter := bson.M{"key": key}

	type Result struct {
		Value string `bson:"value"`
	}

	var result Result
	collection.FindOne(context.Background(), filter).Decode(&result)
	fmt.Println(result.Value)
	return result.Value, nil

}

//Set data to mongo
func (store *Store) Set(key, value string) error {
	type saveStruct struct {
		Key   string
		Value string
	}
	newSaveStruct := saveStruct{Key: key, Value: value}

	_, err := collection.InsertOne(context.Background(), newSaveStruct)
	if err != nil {
		return err
	}

	return nil
}
