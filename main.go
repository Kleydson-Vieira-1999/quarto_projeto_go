package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kleydson-Vieira-1999/quarto_projeto/database"
	"github.com/Kleydson-Vieira-1999/quarto_projeto/routes"
	"github.com/gorilla/handlers"
)

func main() {
	database.ConnectDB()
	fmt.Println("Iniciando Servidor Rest 8080")
	router := routes.Setup()
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
