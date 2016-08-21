package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type users struct{}

type S01Form struct {
	UserName string `form:"userName" validate:"min=2,max=20"`
	Password string `form:"password" validate:"min=4,max=20"`
}

func ValidateUser(c *gin.Context) (*S01Form, error) {
	obj := &S01Form{}
	c.Bind(obj)

	return obj, validator.Validate(obj)
}