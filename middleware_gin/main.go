package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/log", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(200, name)
	})

	r.GET("/panic", func(c *gin.Context) {
		panic("panic test")
	})

	r.Run()
}
