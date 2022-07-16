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
	routes := routesAuthentication
	routes = append(routes, routesUsers...)

	for _, route := range routes {

		if route.WithAuthentication {
			r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
		}

		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
