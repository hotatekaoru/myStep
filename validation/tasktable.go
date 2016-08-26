package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

var validate *validator.Validate

type S42Form struct {
	New         bool    `form:"new"`
	TypeId      int     `form:"typeId" validate:"required,gte=1,lte=3"`
	ContentName string  `form:"contentName" validate:"required"`
	Point       float64 `form:"point" validate:"required,lte=10"`
	Par         int     `form:"par" validate:"required,gte=1,lte=100"`
	UnitId      int     `form:"unitId" validate:"required,gte=1,lte=2"`
}

func ValidateTask(c *gin.Context) (*S42Form, error) {
	config := &validator.Config{TagName: "validate"}

	validate = validator.New(config)

	println("typeId => " + c.PostForm("typeId"))
	println("contentName => " + c.PostForm("contentName"))
	println("point => " + c.PostForm("point"))
	println("unitId => " + c.PostForm("unitId"))
	println("new => " + c.PostForm("new"))
	obj := &S42Form{}
	c.Bind(obj)


	return obj, validate.Struct(obj)
}