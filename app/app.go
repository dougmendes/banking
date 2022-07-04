package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("SERVER_ADDRESS and SERVER_PORT must be set")
	}
}
func Start() {

	sanityCheck()
	router := mux.NewRouter()

	//wiring
	//IN THE MOMENT APPLICATION WILL START IT WILL WIRE THE COMPLETE APPLICATION AND THE HEXAGONAL ARCHITETURE WILL
	//BE APPLIED IN MY APPLICATION
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	//starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
