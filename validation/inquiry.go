package validation

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"myStep/constant"
	"time"
)

type S21Form struct {
	UserCheck []int  `form:"userCheck"`
	DateFrom  string `form:"dateFrom"`
	DateEnd   string `form:"dateEnd"`
	TypeCheck []int  `form:"typeCheck"`
}

func ValidateS21Form(c *gin.Context) (*S21Form, error) {
	validate = validator.New()
	validate.RegisterStructValidation(S21FormValidation, S21Form{})

	obj := &S21Form{}
	c.Bind(obj)

	return obj, validate.Struct(obj)
}

func S21FormValidation(sl validator.StructLevel) {
	form := sl.Current().Interface().(S21Form)

	// 日付チェック
	if form.DateFrom != "" {
		_, ok := time.Parse(constant.DATE_FORMAT_YYYYMMDD, form.DateFrom)
		if ok != nil {
			sl.ReportError(form, "DateFrom", "dateFrom", form.DateFrom, "")
			return
		}
	}
	if form.DateEnd != "" {
		_, ok := time.Parse(constant.DATE_FORMAT_YYYYMMDD, form.DateEnd)
		if ok != nil {
			sl.ReportError(form, "DateEnd", "dateEnd", form.DateEnd, "")
			return
		}
	}
}
