package main

import (
	"api/src/configuration"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configuration.Load()

	fmt.Println("Starting server on port", configuration.ApiPort)

	r := router.Generate()

	fmt.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.ApiPort), r))
}
