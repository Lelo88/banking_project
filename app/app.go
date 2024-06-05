package main

import (
	"log"
	"net/http"
)

func Start(){
	// definimos las rutas
	http.HandleFunc("/greet", greet)
	//definimos una nueva ruta para obtener todos los customers
	http.HandleFunc("/customers", getAllCustomers)
	
	// Empezamos el servidor
	log.Fatal(http.ListenAndServe(":8000", nil))
}