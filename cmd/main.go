package main

import (
	"github.com/Elias2389/api-rest/handler"
	"github.com/Elias2389/api-rest/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	log.Println("Init server in port: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}
