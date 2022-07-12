package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesAuthentication = []Route{
	{
		Uri:                "/login",
		Method:             http.MethodPost,
		Handler:            controllers.Login,
		WithAuthentication: false,
	},
}
