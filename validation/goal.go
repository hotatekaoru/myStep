package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type S32Form struct {
	New       bool   `form:"new"`
	Year      string `form:"year"`
	Month     string `form:"month"`
	Coding    int    `form:"coding"`
	Training  int    `form:"training"`
	HouseWork int    `form:"housework"`
	UserId    int
}

func ValidateS32Form(c *gin.Context) (*S32Form, error) {
	validate = validator.New()
	obj := &S32Form{}
	c.Bind(obj)
	return obj, validate.Struct(obj)
}
