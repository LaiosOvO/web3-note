package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/someJson", func(c *gin.Context) {

		name := []string{"iena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, name)
	})

	r.Run(":8080")
}
