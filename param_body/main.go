package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, "")
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(body))
	})
	r.Run()
}
