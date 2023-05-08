package app

import "server2/models"

var Module models.Module = models.Module{
	Name:      "App",
	Prefix:    "/",
	Endpoints: Endpoints,
}

var Endpoints = []models.Endpoint{
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
