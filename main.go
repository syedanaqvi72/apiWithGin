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

	r.Run(":8080")
}

//inatallation of Gin "go get -u github.com/gin-gonic/gin"

//I have error handling for the POST request. If there's an issue with the JSON data or binding, it will return a 400 Bad Request response with an error message. Make sure your JSON data is correctly formatted.
//
