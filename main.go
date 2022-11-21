package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

//In struct first letter capitalized for export

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = map[string]Customer{
	"1": {ID: 1, Name: "Antonino", Role: "Customer", Email: "test@test.com", Phone: 3333333333, Contacted: true},
	"2": {ID: 2, Name: "Pietro", Role: "Customer", Email: "test1@test.com", Phone: 3333333332, Contacted: false},
	"3": {ID: 3, Name: "Giuseppe", Role: "Customer", Email: "test2@test.com", Phone: 3333333331, Contacted: true},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(customers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", string(res))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	if value, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		res, err := json.Marshal(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	body, readErr := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if readErr != nil {
		fmt.Println(readErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := customers[strconv.Itoa(customer.ID)]; ok {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if customer.ID == 0 {
		customer.ID = rand.Intn(100)
	}
	customers[strconv.Itoa(customer.ID)] = customer
	res, err := json.Marshal(customers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(res))
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer Customer
	id := params["id"]
	body, readErr := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if readErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		customers[id] = customer
		res, err := json.Marshal(customers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		delete(customers, id)
		res, err := json.Marshal(customers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(res))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers", addCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
