package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ipAuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		clientIP := c.ClientIP()
		flag := false
		for _, host := range ipList {
			if host == clientIP{
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, "%s not in ip list", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(ipAuthMiddleware())
	r.GET("/test", func(c *gin.Context){
		c.String(200, "hello test")
	})

	r.Run()
}