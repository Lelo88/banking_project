package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Definimos un struct para utilizar de ejemplo con json
type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func getAllCustomers (w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", Zipcode: "10001"},
		{Name: "Jane Doe", City: "New York", Zipcode: "10001"},
		{Name: "Jack Doe", City: "New York", Zipcode: "10001"},
	}

	// vamos a preguntar si el contentype es xml o json, para poder responder con el formato correcto.
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}