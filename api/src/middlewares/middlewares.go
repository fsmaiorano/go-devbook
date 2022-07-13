package middlewares

import (
	"api/src/helpers"
	"api/src/security"
	"log"
	"net/http"
)

// Logging the request routes
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		// Do something before the request
		next.ServeHTTP(w, r)
		// Do something after the request
	}
}

// Authentication verifies if the user has a valid token
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			helpers.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
