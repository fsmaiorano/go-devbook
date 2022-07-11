package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response is a helper function that returns a JSON response
func Json(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal("Error encoding data: " + err.Error())
	}
}

// Response is a helper function that returns a error response
func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
