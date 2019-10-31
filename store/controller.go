package store

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

func (c *Controller) AddStore(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Bad Data"))
		log.Println(err)
		return
	}
	r.Body.Close()
	var store Store
	json.Unmarshal(body, &store)
	insertedID, err := c.Repository.AddStore(store)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	respBody := AddResponse{
		NewId: insertedID,
	}
	json.NewEncoder(w).Encode(respBody)
}

func (c *Controller) GetStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	store, err := c.Repository.GetOneStore(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store)
}

func (c *Controller) UpdateStore(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Bad Data"))
		log.Println(err)
		return
	}
	defer r.Body.Close()

	var store Store
	json.Unmarshal(body, &store)

	vars := mux.Vars(r)
	id := vars["id"]
	err = c.Repository.UpdateStore(id, store)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}