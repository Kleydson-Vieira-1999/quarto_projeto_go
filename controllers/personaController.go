package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kleydson-Vieira-1999/quarto_projeto/database"
	"github.com/Kleydson-Vieira-1999/quarto_projeto/models"
	"github.com/gorilla/mux"
)

type PersonaController struct {
}

func NewPersonaController() *PersonaController {
	return &PersonaController{}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func (pc *PersonaController) ListPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func (pc *PersonaController) GetPersonaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personaId := vars["id"]
	var p models.Personality
	database.DB.First(&p, personaId)
	json.NewEncoder(w).Encode(p)
}

func (pc *PersonaController) CreatePersona(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Personality
	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}

func (pc *PersonaController) EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personaId := vars["id"]
	var personaEdit models.Personality

	database.DB.First(&personaEdit, personaId)
	json.NewDecoder(r.Body).Decode(&personaEdit)

	database.DB.Save(&personaEdit)
	json.NewEncoder(w).Encode(personaEdit)
}

func (pc *PersonaController) DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personaId := vars["id"]
	var persona models.Personality
	// database.DB.First(&persona, personaId)
	database.DB.Delete(&persona, personaId)
	json.NewEncoder(w).Encode(persona)
}
