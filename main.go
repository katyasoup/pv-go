package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Port string
}

func main() {
	setUpRoutes()
}

func setUpRoutes() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,
			"pong",
		)
	})
	r.Run() // Default is :8080
}
