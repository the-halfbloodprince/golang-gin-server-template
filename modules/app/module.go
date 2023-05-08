package app

import "server2/models"

// GetEndpoints returns all the endpoints for a group/module in the app
func GetEndpoints() (string, []models.Endpoint) {

	prefix := "/"

	Endpoints := []models.Endpoint{
		{
			Method:         "GET",
			Path:           "/",
			ControllerFunc: Home,
		},
		{
			Method:         "GET",
			Path:           "/about",
			ControllerFunc: About,
		},
		{
			Method:         "POST",
			Path:           "/hi",
			ControllerFunc: Hi,
		},
	}

	return prefix, Endpoints

}
