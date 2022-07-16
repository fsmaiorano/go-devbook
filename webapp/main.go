package main

import (
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	println("Hello World")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":3000", r))
}
