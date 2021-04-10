package routers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/customers/services"
	"github.com/customers/entities"
)


func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := services.GetCustomers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	customer := entities.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatal(err)
	}
	services.CreateCustomer(customer)
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(customer)
}

func CustomersRouter(r *mux.Router) {
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
}
