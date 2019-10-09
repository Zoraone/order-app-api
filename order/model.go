package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id    primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Store string             `json:"store"`
}

type OrderCart struct {
	Username string          `json:"username"`
	OrderId  string          `json:"orderId"`
	Items    []OrderCartItem `json:"items"`
}

type OrderCartItem struct {
	Uuid           string                   `json:"uuid"`
	Title          string                   `json:"title"`
	Price          int                      `json:"price"`
	Customizations []OrderCartCustomization `json:"customizations"`
}

type OrderCartCustomization struct {
	Uuid    string            `json:"uuid"`
	Title   string            `json:"title"`
	Options []OrderCartOption `json:"options"`
}

type OrderCartOption struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
