package main

import (
	"api/src/configuration"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal("Error generating key")
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)

// 	fmt.Println("Generated key:", stringBase64)
// }

func main() {
	configuration.Load()

	r := router.Generate()

	fmt.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.ApiPort), r))
}
