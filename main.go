package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"server2/config"
	"strings"
)

func main() {

	// initialize the router
	r := gin.Default()

	// attach router in AppConfig
	config.AppConfig.Router = r

	// set a func map for the templates
	r.SetFuncMap(template.FuncMap{
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
	})

	// map static paths
	r.Static("/assets", "./static")

	// load view templates
	r.LoadHTMLGlob("templates/*.html")

	// initialize routes
	InitializeRoutes()

	// start the app
	_ = r.Run()

}
