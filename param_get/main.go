package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_last_name")
		c.JSON(http.StatusOK, gin.H{
			"first_name": firstName,
			"last_name":  lastName,
		})
	})
	r.Run(":8080")
}
