package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateSearch",
		"POST",
		"/api/search",
		CreateSearch,
	},
	Route{
		"CreatePublish",
		"POST",
		"/api/publish",
		CreatePublish,
	},
	// Route{
	// 	"GetRoutes"
	// 	"GET"
	// 	"/api/get_routes"
	// 	GetRoutes
	// },
	// Route{
	// 	"Publish",
	// 	"POST",
	// 	"/api/publish",
	// 	Publish,
	// },
	// Route{
	// 	"GetRoute",
	// 	"GET",
	// 	"/api/getroute",
	// 	GetRoute,
	// },
}
