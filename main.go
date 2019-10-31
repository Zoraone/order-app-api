package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Zoraone/order-app-api/item"
	"github.com/Zoraone/order-app-api/order"
	"github.com/Zoraone/order-app-api/store"

	"github.com/gorilla/mux"
)

var orderController = &order.Controller{Repository: order.Repository{}}
var storeController = &store.Controller{Repository: store.Repository{}}
var itemController = &item.Controller{Repository: item.Repository{}}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/api/order/get/", orderController.GetAllOrders).Methods("GET")
	router.HandleFunc("/api/order/get/{id}", orderController.GetOrder).Methods("GET")
	router.HandleFunc("/api/order/add/", orderController.AddOrder).Methods("POST")
	router.HandleFunc("/api/order/carts/get/{id}", orderController.GetCartsInOrder).Methods("GET")

	router.HandleFunc("/api/cart/add/", orderController.AddOrderCart).Methods("POST")

	router.HandleFunc("/api/store/add/", storeController.AddStore).Methods("POST")
	router.HandleFunc("/api/store/get/{id}", storeController.GetStore).Methods("GET")
	router.HandleFunc("/api/store/update/{id}", storeController.UpdateStore).Methods("POST")

	router.HandleFunc("/api/item/add/", itemController.AddItem).Methods("POST")
	router.HandleFunc("/api/item/get/{id}", itemController.GetItem).Methods("GET")
	router.HandleFunc("/api/item/update/{id}", itemController.UpdateItem).Methods("POST")

	log.Fatal(http.ListenAndServe(getPort(), corsMiddleware(router)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8080" // Local default
}
