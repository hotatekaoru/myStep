package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type S32Form struct {
	New       bool   `form:"new"`
	Year      string `form:"year" validate:"required"`
	Month     string `form:"month" validate:"required"`
	Coding    int    `form:"coding" validate:"required"`
	Training  int    `form:"training" validate:"required"`
	HouseWork int    `form:"housework" validate:"required"`
	Continue  int    `form:"continue"`
	UserId    int
}

func ValidateS32Form(c *gin.Context) (*S32Form, error) {
	validate = validator.New()
	obj := &S32Form{}
	c.Bind(obj)
	return obj, validate.Struct(obj)
}
