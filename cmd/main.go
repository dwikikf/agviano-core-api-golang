package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Application entry point
	// ping check
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on
}
