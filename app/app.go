package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(){

	mux := mux.NewRouter()

	// definimos las rutas
	mux.HandleFunc("/greet", greet)
	//definimos una nueva ruta para obtener todos los customers
	mux.HandleFunc("/customers", getAllCustomers)
	
	// Empezamos el servidor
	log.Fatal(http.ListenAndServe(":8000", nil))
}