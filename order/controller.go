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
	orders, err := c.Repository.GetAllOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order, err := c.Repository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (c *Controller) AddOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("500 - Bad Data"))
		log.Println(err)
		return
	}
	r.Body.Close()
	var order Order
	json.Unmarshal(body, &order)
	insertedID, err := c.Repository.AddOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	respBody := AddOrderResponse{
		NewId: insertedID,
	}
	json.NewEncoder(w).Encode(respBody)
}

func (c *Controller) AddOrderCart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("500 - Bad Data"))
		log.Println(err)
		return
	}
	r.Body.Close()
	var cart OrderCart
	json.Unmarshal(body, &cart)
	insertedID, err := c.Repository.AddOrderCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	respBody := AddCartResponse{
		NewId: insertedID,
	}
	json.NewEncoder(w).Encode(respBody)
}

func (c *Controller) GetCartsInOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	carts, err := c.Repository.GetCartsInOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carts)
}
