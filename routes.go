package main

import (
	"server2/config"
	"server2/models"
	"server2/modules/app"
)

// InitializeRoutes initializes all the routes
func InitializeRoutes() {

	// MANUALLY ADD ROUTE GROUPS HERE
	// app controllers
	initializeRouteGroup(app.GetEndpoints())
}

// initializeRouteGroup initializes a route group
func initializeRouteGroup(prefix string, endpoints []models.Endpoint) {

	// set up the group router
	subgroupRouter := config.AppConfig.Router.Group(prefix)

	for _, endpoint := range endpoints {

		// set the allowed methods and their corresponding handlers in the framework
		switch endpoint.Method {
		case "GET":
			{
				subgroupRouter.GET(endpoint.Path, endpoint.ControllerFunc)
			}
		case "POST":
			{
				subgroupRouter.POST(endpoint.Path, endpoint.ControllerFunc)
			}
		case "PUT":
			{
				subgroupRouter.PUT(endpoint.Path, endpoint.ControllerFunc)
			}
		case "DELETE":
			{
				subgroupRouter.DELETE(endpoint.Path, endpoint.ControllerFunc)
			}
		default:
			{
				panic("Undefined method route encountered")
			}
		}
	}
}
