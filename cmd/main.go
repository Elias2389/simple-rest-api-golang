package main

import (
	"github.com/Elias2389/api-rest/authorization"
	"github.com/Elias2389/api-rest/handler"
	"github.com/Elias2389/api-rest/storage"
	"log"
	"net/http"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Certificates not found: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Init server in port: 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}
