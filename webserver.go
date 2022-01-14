package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startWebServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.POST("/api/barcode/:ean/info", func(c *gin.Context) {
		barcode := Barcode{EAN: c.Param("ean")}
		product := barcode.Scaned()

		c.JSON(200, product)
	})

	r.POST("/api/barcode/:ean/subtract/:amount", func(c *gin.Context) {
		barcode := Barcode{EAN: c.Param("ean")}
		product := barcode.Scaned()

		amount, _ := strconv.Atoi(c.Param("amount"))
		fmt.Println(amount)
		product.Subtract(amount)

		c.JSON(200, product)
	})

	r.POST("/api/barcode/:ean/add/:amount", func(c *gin.Context) {
		barcode := Barcode{EAN: c.Param("ean")}
		product := barcode.Scaned()

		amount, _ := strconv.Atoi(c.Param("amount"))
		fmt.Println(amount)
		product.Add(amount)

		c.JSON(200, product)
	})

	r.Run()
}
