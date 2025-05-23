package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{
			Id:      1,
			Name:    "Albert Einstein",
			History: "Albert Einstein was a theoretical physicist who developed the theory of relativity, one of the two pillars of modern physics (alongside quantum mechanics).",
		},
		{
			Id:      2,
			Name:    "Marie Curie",
			History: "Marie Curie was a Polish-born physicist and chemist who conducted pioneering research on radioactivity.",
		},
	}

	fmt.Println("Initializing the rest api server...")
	routes.HandleRequest()
}
