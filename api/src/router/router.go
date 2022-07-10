package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate application routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}
