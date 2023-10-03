package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin router.
	r := gin.Default()

	// Define a route that responds to GET requests.
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Define a route that responds to POST requests.
	r.POST("/postdata", func(c *gin.Context) {
		var requestData struct {
			Data string `json:"data"`
		}

		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"received_data": requestData.Data})
	})

	// Start the web server on port 8080.
	r.Run(":8080")
}
