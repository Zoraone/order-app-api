package item

import (
	"context"

	"github.com/Zoraone/order-app-api/util"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {}

const ITEM_COLLECTION = "items"

func (r Repository) AddItem(item Item) (interface{}, error) {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(ITEM_COLLECTION)
	defer client.Disconnect(context.Background())
	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func (r Repository) GetItem(id string) (Item, error) {
	client := util.GetClient()
	var item Item
	collection := client.Database(util.GetDBName()).Collection(ITEM_COLLECTION)
	defer client.Disconnect(context.Background())
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r Repository) UpdateItem(id string, item Item) error {
	client := util.GetClient()
	defer client.Disconnect(context.Background())
	collection := client.Database(util.GetDBName()).Collection(ITEM_COLLECTION)
	filter := bson.D{{"_id", id}}
	update := bson.D{{Key: "$set", Value: item}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}