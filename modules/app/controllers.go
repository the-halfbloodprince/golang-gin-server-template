package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// controllers here

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is the home page content. wooooo",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"content": "This is the about page content",
	})
}

func Hi(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"content": "Hi everyone",
	})
}
