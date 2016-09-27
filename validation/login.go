package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type S01Form struct {
	UserName string `form:"userName" validate:"required,gte=2,lte=20"`
	Password string `form:"password" validate:"required,gte=4,lte=20"`
}

type J01Form struct {
	UserName string `form:"userName" json:"userName" validate:"required,gte=2,lte=20"`
	Password string `form:"password" json:"password" validate:"required,gte=4,lte=20"`
}

func ValidateS01Form(c *gin.Context) (*S01Form, error) {
	validate = validator.New()

	obj := &S01Form{}
	c.Bind(obj)

	err := validate.Struct(obj)

	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"userName": c.PostForm("userName"),
			"error": []error{
				err,
			},
		})
	}
	return obj, err
}

func ValidateS01FormFromJSON(c *gin.Context) (*S01Form, error) {
	validate = validator.New()

	j01 := &J01Form{}
	s01 := &S01Form{}
	c.Bind(j01)

	err := validate.Struct(j01)

	if err != nil {
		return s01, err
	}

	s01.UserName = j01.UserName
	s01.Password = j01.Password

	return s01, err
}
