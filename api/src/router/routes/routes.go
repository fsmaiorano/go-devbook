package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, routesAuthentication...)

	for _, route := range routes {

		if route.WithAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authentication(route.Handler))).Methods(route.Method)
		}

		r.HandleFunc(route.Uri, middlewares.Logger(route.Handler)).Methods(route.Method)
	}

	return r
}
