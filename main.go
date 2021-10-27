package main

import (
	"net/http"

	"github.com/andkolbe/gin-docker-blog/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // initialize Gin router

	router.GET("/", func(c *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})
	router.Run(":8000")
}