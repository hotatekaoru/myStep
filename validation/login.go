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

func ValidateUser(c *gin.Context) (*S01Form, error) {
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
