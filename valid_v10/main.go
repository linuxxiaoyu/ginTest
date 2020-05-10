package main

import (
	"github.com/gin-gonic/gin"
	en1 "github.com/go-playground/locales/en"
	zh1 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// Person defines the information of a person.
type Person struct {
	Age     int    `form:"age" validate:"required,gt=11"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address"`
}

func main() {
	Validate := validator.New()
	zh := zh1.New()
	en := en1.New()
	Uni := ut.New(zh, en)

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			return
		}

		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			return
		}

		c.String(200, "%v", person)
	})

	r.Run(":8080")
}
