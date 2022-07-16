package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesAuthentication = []Route{
	{
		Uri:                "/",
		Method:             http.MethodGet,
		Handler:            controllers.Authentication,
		WithAuthentication: false,
	},
	{
		Uri:                "/login",
		Method:             http.MethodGet,
		Handler:            controllers.Authentication,
		WithAuthentication: false,
	},
}
