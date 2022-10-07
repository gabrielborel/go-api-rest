package controllers

import (
	"api-rest/database"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality

	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality

	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	
	database.DB.First(&personality, id)
	
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
