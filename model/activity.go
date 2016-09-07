package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"myStep/constant"
	"myStep/validation"
	"net/http"
	"time"
)

type Activity struct {
	GormModel
	UserId      int       `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	TypeId      int       `gorm:"not null"`
	TaskId      int       `gorm:"not null"`
	Point       float64   `gorm:"not null"`
	WorkingTime int
	Comment     string `gorm:"varchar(300)"`
}

type task struct {
	TaskId  int
	Content string
}

type S11Form struct {
	New           bool
	UserId        int
	Date          string
	TypeId        int
	TaskId        int
	CodingList    []task
	TrainingList  []task
	HouseworkList []task
}

type S12Form struct {
	New            bool
	UserName       string
	Date           string
	TypeName       string
	Content        string
	Point          float64
	WorkingTime    int
	UnitId         int
	Comment        string
	PointParMinute float64
}

type s21Activity struct{
	ActivityId     int
	UserName       string
	Date           string
	TypeName       string
	Content        string
	UnitId         int
	WorkingTime    int
	Point          float64
}

type S21Form struct {
	Activity       []s21Activity
}

var userMap = map[int]string{1: "Kaoru", 2: "Yuri"}

/* DB操作 */
func selectActivity() *[]Activity{
	activity := []Activity{}
	db.Debug().Model(&Activity{}).Order("id").Find(&activity)
	return &activity

}

func CreateActivity(s11 *validation.S11Form, s12 *validation.S12Form) {
	date, _ := time.Parse(constant.DATE_FORMAT_YYYYMMDD, s11.Date)
	activity := Activity{
		UserId:      s11.UserId,
		Date:        date,
		TypeId:      s11.TypeId,
		TaskId:      s11.TaskId,
		Point:       s12.Point,
		WorkingTime: s12.WorkingTime,
		Comment:     s12.Comment,
	}

	db.Debug().Create(&activity)
}

/* form（外部出力）操作 */
func GetS11FormRegister(typeId int) *S11Form {

	form := S11Form{
		New:    true,
		UserId: 1,
		Date:   time.Now().Format(constant.DATE_FORMAT_YYYYMMDD),
		TypeId: typeId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	return &form
}

func GetS11FormBySession(s validation.S11Form) *S11Form {

	form := S11Form{
		New:    s.New,
		UserId: s.UserId,
		Date:   s.Date,
		TypeId: s.TypeId,
		TaskId: s.TaskId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	return &form
}

func SetTaskToS11Form(form *S11Form, taskList *[]Task) {
	for _, v := range *taskList {
		t := task{
			TaskId:  v.Id,
			Content: v.Content,
		}
		switch v.TypeId {
		case 1:
			form.CodingList = append(form.CodingList, t)
		case 2:
			form.TrainingList = append(form.TrainingList, t)
		case 3:
			form.HouseworkList = append(form.HouseworkList, t)
		}
	}
}

func ReturnS11InputErr(input *validation.S11Form, errs error, c *gin.Context) {
	form := S11Form{
		New:    input.New,
		UserId: input.UserId,
		Date:   input.Date,
		TypeId: input.TypeId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field()+constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "activity_register1.html", gin.H{
		"errorList": err,
		"form":      form,
	})
}

func ReturnS12InputErr(s11form *validation.S11Form, input *validation.S12Form, errs error, c *gin.Context) {
	userName, _ := userMap[s11form.UserId]
	typeName, _ := typeMap[s11form.TypeId]
	task := SelectTaskById(s11form.TaskId)

	form := S12Form{
		New:            input.New,
		UserName:       userName,
		Date:           s11form.Date,
		TypeName:       typeName,
		Content:        task.Content,
		Point:          input.Point,
		WorkingTime:    input.WorkingTime,
		UnitId:         task.UnitId,
		Comment:        input.Comment,
		PointParMinute: input.Point / 10,
	}

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field()+constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "activity_register2.html", gin.H{
		"errorList": err,
		"form":      form,
	})
}

func GetS12FormRegister(input *validation.S11Form) *S12Form {
	userName, _ := userMap[input.UserId]
	typeName, _ := typeMap[input.TypeId]
	task := SelectTaskById(input.TaskId)

	form := S12Form{
		New:            true,
		UserName:       userName,
		Date:           input.Date,
		TypeName:       typeName,
		Content:        task.Content,
		Point:          task.Point,
		WorkingTime:    10,
		UnitId:         task.UnitId,
		Comment:        "",
		PointParMinute: task.Point / 10,
	}

	return &form
}

func GetS21Form() *S21Form {
	form := S21Form{}

	activities := selectActivity()
	form.Activity = convertActivitiesToS21Form(activities)

	return &form
}

func convertActivitiesToS21Form(activities *[]Activity) []s21Activity {
	var s21acts []s21Activity
	tasks := SelectAllTask()
	for _, activity := range *activities {
		s21acts = append(s21acts, convertOneActivityToS21Form(&activity, tasks))
	}
	return s21acts
}

func convertOneActivityToS21Form(activity *Activity, tasks *[]Task) s21Activity {
	form := s21Activity{}

	form.ActivityId = activity.Id

	s, _ := userMap[activity.UserId]
	form.UserName = s

	form.Date = activity.Date.Format(constant.DATE_FORMAT_YYYYMMDD)

	s, _ = typeMap[activity.TypeId]
	form.TypeName = s

	form.WorkingTime = activity.WorkingTime

	form.Point = activity.Point

	for _, v := range *tasks {
		if v.Id == activity.TaskId {
			form.Content = v.Content
			form.UnitId = v.UnitId
		}
	}

	return form
}