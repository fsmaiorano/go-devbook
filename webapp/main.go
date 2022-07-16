package main

import (
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	println("Hello World")

	utils.LoadTemplates()
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":3000", r))
}
