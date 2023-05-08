package config

import "github.com/gin-gonic/gin"

// structure of the AppConfig
type appConfig struct {
	Router *gin.Engine
}

// AppConfig is the global config for the app to be available everywhere in the app
var AppConfig appConfig = appConfig{}
