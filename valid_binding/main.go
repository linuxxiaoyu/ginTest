package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     int    `form:"age" binding:"required,gt=11"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusOK, "%v", err)
		} else {
			c.String(http.StatusOK, "%v", person)
		}
	})
	r.Run()
}
