package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.GetPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
