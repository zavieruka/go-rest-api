package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
