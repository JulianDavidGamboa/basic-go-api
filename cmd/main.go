package main

import (
	"log"
	"net/http"

	"github.com/JulianDavidGamboa/basic-go-api/cmd/authorization"
	"github.com/JulianDavidGamboa/basic-go-api/handler"
	"github.com/JulianDavidGamboa/basic-go-api/storage"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
	//calle 15 a # 101 - 60 barrio ciudad jardin

}
