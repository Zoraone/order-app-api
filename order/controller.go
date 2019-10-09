package order

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders := c.Repository.GetAllOrders()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)

	return
}

func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order := c.Repository.GetOrder(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)

	return
}

func (c *Controller) AddOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("Error AddOrder", err)
	}
	r.Body.Close()
	var order Order
	json.Unmarshal(body, &order)
	c.Repository.AddOrder(order)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) AddOrderCart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("Error AddOrderCart", err)
	}
	r.Body.Close()
	var cart OrderCart
	json.Unmarshal(body, &cart)
	c.Repository.AddOrderCart(cart)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) GetCartsInOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	carts := c.Repository.GetCartsInOrder(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
}