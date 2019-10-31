package store

import (
	"context"

	"github.com/Zoraone/order-app-api/util"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct{}

const STORE_COLLECTION = "stores"

func (r Repository) AddStore(store Store) (interface{}, error) {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(STORE_COLLECTION)
	defer client.Disconnect(context.Background())
	insertResult, err := collection.InsertOne(context.TODO(), store)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func (r Repository) GetOneStore(id string) (Store, error) {
	client := util.GetClient()
	var store Store
	collection := client.Database(util.GetDBName()).Collection(STORE_COLLECTION)
	defer client.Disconnect(context.Background())
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&store)
	if err != nil {
		return Store{}, err
	}
	return store, nil
}

func (r Repository) UpdateStore(id string, store Store) error {
	client := util.GetClient()
	defer client.Disconnect(context.Background())
	collection := client.Database(util.GetDBName()).Collection(STORE_COLLECTION)
	filter := bson.D{{"_id", id}}
	update := bson.D{{Key: "$set", Value: store}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}