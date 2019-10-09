package store

import (
	"context"
	"log"

	"../util"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct{}

const STORE_COLLECTION = "stores"

func (r Repository) AddStore(store Store) interface{} {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(STORE_COLLECTION)
	defer client.Disconnect(context.Background())
	insertResult, err := collection.InsertOne(context.TODO(), store)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID
}

func (r Repository) GetOneStore(id string) Store {
	client := util.GetClient()
	var store Store
	collection := client.Database(util.GetDBName()).Collection(STORE_COLLECTION)
	defer client.Disconnect(context.Background())
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	return store
}
