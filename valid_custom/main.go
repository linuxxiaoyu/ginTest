package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookableDate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(time.Time); ok {
		return v.Unix() > time.Now().Unix()
	}
	return false
}

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookableDate", bookableDate)
	}

	r.GET("/booking", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.String(http.StatusInternalServerError, "%v", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok!",
				"b":       b,
			})
		}
	})

	r.Run()
}
