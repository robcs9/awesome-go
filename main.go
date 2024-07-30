package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func main() {
	
	// Instanciating Gin middleware (?)
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	// HTTP routing endpoints
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Awesome",
		})
	})
	
	// Run server
	e.Run(":3000")

}

func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
			"name": "Awesome",
		},
	)
}