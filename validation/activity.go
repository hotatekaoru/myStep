package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"errors"
	"myStep/constant"
)

type S11Form struct {
	New         bool      `form:"new"`
	UserId      int       `form:"userId" validate:"required,gte=1,lte=2"`
	//Date       string    `form:"date" validate:"required"`
	TypeId      int       `form:"typeId" validate:"required,gte=1,lte=3"`
	ContentId   int
	ContentId1  int       `form:"contentId1"`
	ContentId2  int       `form:"contentId2"`
	ContentId3  int       `form:"contentId3"`
}

type S12Form struct {
	New         bool      `form:"new"`
	TaskId      int       `form:"taskId"`
	ContentName string    `form:"contentName" validate:"required"`
	Point       float64   `form:"point" validate:"required,lte=10"`
	Par         int       `form:"workingTime" validate:"required,gte=1,lte=100"`
	UnitId      int       `form:"unitId" validate:"required,gte=1,lte=2"`
	Comment     string    `form:"comment" validate:"lte=200"`
}

func ValidateS11Form(c *gin.Context) (*S11Form, error) {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)

	obj := &S11Form{}
	c.Bind(obj)
	return obj, validate.Struct(obj)
}

func ValidateS12Form(c *gin.Context) (*S12Form, error) {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)

	obj := &S12Form{}
	c.Bind(obj)

	return obj, validate.Struct(obj)
}

func ValidateContentId(input *S11Form) bool {
	switch input.TypeId {
	case 1:
		if input.ContentId1 == 0 {
			return false
		}
		input.ContentId = input.ContentId1
	case 2:
		if input.ContentId2 == 0 {
			return false
		}
		input.ContentId = input.ContentId2
	case 3:
		if input.ContentId3 == 0 {
			return false
		}
		input.ContentId = input.ContentId3
	default:
		// validationで除外済みだけど
		return false
	}
	return true
}