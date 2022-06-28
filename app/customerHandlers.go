package app

import (
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// CustomerHandler CREATED A CUSTOMER HANDLER AND DEFINED A CUSTOMER SERVICE WICH WILL BEA DEPENDECY TO THIS HANDLERS
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandler) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customers, err := ch.service.GetCustomersById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		writeResponse(w, http.StatusOK, customers)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
