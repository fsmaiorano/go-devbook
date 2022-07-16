package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	{
		Uri:                "/posts",
		Method:             http.MethodPost,
		Handler:            controllers.CreatePost,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts",
		Method:             http.MethodGet,
		Handler:            controllers.GetPosts,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodGet,
		Handler:            controllers.GetPost,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodPut,
		Handler:            controllers.UpdatePost,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodDelete,
		Handler:            controllers.DeletePost,
		WithAuthentication: true,
	},
	{
		Uri:                "/users/{id}/posts",
		Method:             http.MethodGet,
		Handler:            controllers.FindPostsByUser,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts/{id}/like",
		Method:             http.MethodPost,
		Handler:            controllers.LikePost,
		WithAuthentication: true,
	},
	{
		Uri:                "/posts/{id}/unlike",
		Method:             http.MethodPost,
		Handler:            controllers.UnlikePost,
		WithAuthentication: true,
	},
}
