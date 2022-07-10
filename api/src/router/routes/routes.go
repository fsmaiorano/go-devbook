package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes is a struct that contains all the routes for the application
type Route struct {
	Uri                string
	Method             string
	Handler            func(w http.ResponseWriter, r *http.Request)
	WithAuthentication bool
}

// ConfigureRoutes is a function that configures the routes for the application
func Configuration(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	return r
}
