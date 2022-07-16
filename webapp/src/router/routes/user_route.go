package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesUsers = []Route{
	{
		Uri:                "/signup",
		Method:             http.MethodGet,
		Handler:            controllers.LoadCreateUserPage,
		WithAuthentication: false,
	},
	{
		Uri:                "/users",
		Method:             http.MethodPost,
		Handler:            controllers.CreateUser,
		WithAuthentication: false,
	},
}
