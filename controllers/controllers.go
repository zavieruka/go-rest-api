package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if fmt.Sprintf("%d", personality.Id) == id {
			json.NewEncoder(w).Encode(personality)
			return
		}
	}

	http.Error(w, "Personality not found", http.StatusNotFound)
}
