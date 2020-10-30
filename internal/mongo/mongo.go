// у этого класса должен быть клиенит
package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// var

// type Client struct {

// }
const keyFieldName string = "key"
const valueParam string = "value"

type MongoStore struct {
	Client *mongo.Client
}

var collection *mongo.Collection

//Get data from mongo
func (store *MongoStore) Get(key string) (string, error) {
	filter := bson.M{"key": key}

	var resultString struct {
		result string
	}

	collection.FindOne(context.Background(), filter).Decode(&resultString)
	return resultString.result, nil

}

//Set data to mongo
func (store *MongoStore) Set(key, value string) error {
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
