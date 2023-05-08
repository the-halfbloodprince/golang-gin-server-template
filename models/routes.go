package models

import "github.com/gin-gonic/gin"

// Endpoint is the endpoint structure
type Endpoint struct {
	Method         string
	Path           string
	ControllerFunc gin.HandlerFunc
}
