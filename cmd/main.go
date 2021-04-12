package main

import (
	"github.com/Elias2389/api-rest/authorization"
	"github.com/Elias2389/api-rest/handler"
	"github.com/Elias2389/api-rest/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Certificates not found: %v", err)
	}

	store := storage.NewMemory()
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Init server in port: 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
