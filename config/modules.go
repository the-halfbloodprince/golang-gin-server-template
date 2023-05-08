package config

import (
	"server2/models"
	"server2/modules/app"
)

// Modules is the list of all active modules in the app
// register the modules here
var Modules []models.Module = []models.Module{
	app.Module,
}
