package main

import (
	"log"
	"net/http"

	"github.com/biFebriansyah/gorest/routers"
)

func main() {
	mainRutes := routers.New()

	log.Println("Serve API on http://localhost:8080")
	err := http.ListenAndServe(":8080", mainRutes)
	if err != nil {
		log.Fatal("Error API")
	}

}
