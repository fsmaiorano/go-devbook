package routes

import (
	"api/src/controllers"
	"net/http"
)

// UserRoutes is a struct that contains all the routes for the user
var routesUsers = []Route{
	{
		Uri:                "/users",
		Method:             http.MethodPost,
		Handler:            controllers.CreateUser,
		WithAuthentication: false,
	},
	{
		Uri:                "/users",
		Method:             http.MethodGet,
		Handler:            controllers.GetUsers,
		WithAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		Handler:            controllers.GetUser,
		WithAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdateUser,
		WithAuthentication: false,
	},
	{
		Uri:                "/users",
		Method:             http.MethodDelete,
		Handler:            controllers.DeleteUser,
		WithAuthentication: false,
	},
}
