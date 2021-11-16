package main

import (
	"github.com/gin-gonic/gin"
)

func startWebServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "pong",
		})
	})
	r.Run()
}
