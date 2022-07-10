package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	r := router.Generate()

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
