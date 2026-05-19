package routes

import (
	"github.com/gorilla/mux"

	"github.com/Kleydson-Vieira-1999/quarto_projeto/controllers"
	"github.com/Kleydson-Vieira-1999/quarto_projeto/middleware"
)

func Setup() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.AddContentTypeMiddleware)

	personaController := controllers.NewPersonaController()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/persona", personaController.ListPersonalities).Methods("GET")
	router.HandleFunc("/api/persona/{id}", personaController.GetPersonaById).Methods("GET")
	router.HandleFunc("/api/persona", personaController.CreatePersona).Methods("POST")
	router.HandleFunc("/api/persona/{id}", personaController.EditPersona).Methods("PUT")
	router.HandleFunc("/api/persona/{id}", personaController.DeletePersona).Methods("DELETE")

	return router
}
