package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"time"
	"myStep/constant"
)


type S11Form struct {
	New         bool      `form:"new"`
	UserId      int       `form:"userId" validate:"required,gte=1,lte=2"`
	Date        string    `form:"date" validate:"required"`
	TypeId      int       `form:"typeId" validate:"required,gte=1,lte=3"`
	TaskId      int
	TaskId1     int       `form:"taskId1"`
	TaskId2     int       `form:"taskId2"`
	TaskId3     int       `form:"taskId3"`
}

type S12Form struct {
	New         bool      `form:"new"`
	WorkingTime int       `form:"workingTime"`
	Point       float64   `form:"point" validate:"required,lte=10"`
	Comment     string    `form:"comment" validate:"lte=300"`
}

func ValidateS11URLQuery(c *gin.Context) int {
	inputTypeId := c.Param("typeId")
	if inputTypeId == "" {
		return 1
	}

	typeId, _ := strconv.Atoi(inputTypeId)
	if 1 <= typeId && typeId <= 3 {
		return typeId
	}
	return 1
}

func ValidateS11Form(c *gin.Context) (*S11Form, error) {
	validate = validator.New()
	validate.RegisterStructValidation(S11FormValidation, S11Form{})

	obj := &S11Form{}
	c.Bind(obj)
	return obj, validate.Struct(obj)
}

func S11FormValidation(sl validator.StructLevel) {

	form := sl.Current().Interface().(S11Form)

	_, ok := time.Parse(constant.DATE_FORMAT_YYYYMMDD, form.Date)
	if ok != nil {
		sl.ReportError(form, "Date", "date", form.Date, "")
		return
	}
}

func SetTaskId(input *S11Form) {
	switch (input.TypeId) {
	case 1:
		input.TaskId = input.TaskId1
	case 2:
		input.TaskId = input.TaskId2
	case 3:
		input.TaskId = input.TaskId3
	}
}

func ValidateS12Form(c *gin.Context) (*S12Form, error) {
	validate = validator.New()

	obj := &S12Form{}
	c.Bind(obj)

	return obj, validate.Struct(obj)
}