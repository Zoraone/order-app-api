package item

import (
	"context"
	"log"

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

func (r Repository) GetItem(id string) Item {
	client := util.GetClient()
	var item Item
	collection := client.Database(util.GetDBName()).Collection(ITEM_COLLECTION)
	defer client.Disconnect(context.Background())
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		log.Fatal(err)
	}
	return item
}