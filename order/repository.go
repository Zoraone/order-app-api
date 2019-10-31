package order

import (
	"context"
	"log"

	"github.com/Zoraone/order-app-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct{}

const ORDER_COLLECTION = "orders"
const CART_COLLECTION = "order_carts"

func (r Repository) GetAllOrders() ([]Order, error) {
	var orders []Order
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(ORDER_COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return orders, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var order Order
		err = cur.Decode(&order)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r Repository) AddOrder(order Order) (interface{}, error) {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(ORDER_COLLECTION)
	defer client.Disconnect(context.Background())
	order.Id = primitive.NewObjectID()
	insertResult, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func (r Repository) GetOrder(id string) (Order, error) {
	client := util.GetClient()
	var order Order
	collection := client.Database(util.GetDBName()).Collection(ORDER_COLLECTION)
	defer client.Disconnect(context.Background())
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	err := collection.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

// Not used
func (r Repository) UpdateOrder(id string, order Order) int64 {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(ORDER_COLLECTION)
	filter := bson.D{{"_id", id}}
	update := bson.D{{Key: "$set", Value: order}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updatedResult.ModifiedCount
}

func (r Repository) DeleteOrder(id string) int64 {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(ORDER_COLLECTION)
	filter := bson.D{{"_id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}

func (r Repository) AddOrderCart(cart OrderCart) (interface{}, error) {
	client := util.GetClient()
	collection := client.Database(util.GetDBName()).Collection(CART_COLLECTION)
	defer client.Disconnect(context.Background())
	insertResult, err := collection.InsertOne(context.TODO(), cart)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func (r Repository) GetCartsInOrder(orderId string) ([]OrderCart, error) {
	client := util.GetClient()

	collection := client.Database(util.GetDBName()).Collection(CART_COLLECTION)
	defer client.Disconnect(context.Background())
	filter := bson.D{{"orderid", orderId}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []OrderCart{}, err
	}
	defer cur.Close(context.TODO())

	var carts []OrderCart
	for cur.Next(context.TODO()) {
		var cart OrderCart
		err = cur.Decode(&cart)
		if err != nil {
			return carts, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}
