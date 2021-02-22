package main

import (
	"log"

	"github.com/JulianDavidGamboa/basic-go-api/cmd/authorization"
	"github.com/JulianDavidGamboa/basic-go-api/handler"
	"github.com/JulianDavidGamboa/basic-go-api/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// mux := http.NewServeMux()

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	// err = http.ListenAndServe(":8080", mux)
	err = e.Start(":8080")

	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
	//calle 15 a # 101 - 60 barrio ciudad jardin

}
