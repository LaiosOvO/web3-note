package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {

		firstName := c.DefaultQuery("firstName", "default value")
		lastName := c.Query("lastName")

		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)

	})

	router.Run(":8080")

}
