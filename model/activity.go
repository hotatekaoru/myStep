package model

import (
	"time"
	"myStep/validation"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"errors"
	"myStep/constant"
)

// 一部タスクテーブルと同じ構成だが、Taskテーブルは変更を許容しているため、
// Activity登録時点のTask情報をActivityテーブルに持つ
type Activity struct {
	GormModel
	UserId      int        `gorm:"not null"`
	Date        time.Time  `gorm:"not null"`
	DateStr     string     `gorm:"-"`
	// ジャンル 1:Coding, 2:Training, 3:Housework
	TypeId      int        `gorm:"not null"`
	ContentId   int        `gorm:"not null"`
	Content     string     `gorm:"not null:varchar(100)"`
	Point       float64    `gorm:"not null"`
	// 1:1回, 2:10分
	UnitId      int        `gorm:"not null"`
	Comment     string     `gorm:"varchar(300)"`
}

type task struct {
	TaskId   int
	Content  string
}

type S11Form struct {
	New              bool
	UserId           int
	Date             string
	TypeId           int
	TaskId           int
	CodingList       []task
	TrainingList     []task
	HouseworkList    []task
}

type S12Form struct {
	New              bool
	UserName         string
	Date             string
	TypeName         string
	Content          string
	Point            float64
	WorkingTime      int
	UnitId           int
	Comment          string
	PointParMinute   float64
}

var userMap = map[int]string{1:"Kaoru", 2:"Yuri"}

/* DB操作 */

/* form（外部出力）操作 */
func GetS11FormRegister(typeId int) *S11Form {

	form := S11Form {
		New:         true,
		UserId:      1,
		Date:        time.Now().Format(constant.DATE_FORMAT_YYYYMMDD),
		TypeId:      typeId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	return &form
}

func GetS11FormBySession(s validation.S11Form) *S11Form {

	form := S11Form {
		New:         s.New,
		UserId:      s.UserId,
		Date:        s.Date,
		TypeId:      s.TypeId,
		TaskId:      s.TaskId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	return &form
}

func SetTaskToS11Form(form *S11Form, taskList *[]Task) {
	for _, v := range *taskList {
		t := task {
			TaskId:   v.Id,
			Content:  v.Content,
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
	form := S11Form {
		New:         input.New,
		UserId:      input.UserId,
		Date:        input.Date,
		TypeId:      input.TypeId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field() + constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "activity_register1.html", gin.H{
		"errorList": err,
		"form": form,
	})
}

func ReturnS12InputErr(s11form *validation.S11Form, input *validation.S12Form, errs error, c *gin.Context) {
	userName, _ := userMap[s11form.UserId]
	typeName, _ := typeMap[s11form.TypeId]
	task := SelectTaskById(s11form.TaskId)

	form := S12Form {
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
		err = append(err, errors.New(v.Field() + constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "activity_register2.html", gin.H{
		"errorList": err,
		"form": form,
	})
}


func GetS12FormRegister(input *validation.S11Form) *S12Form{
	userName, _ := userMap[input.UserId]
	typeName, _ := typeMap[input.TypeId]
	task := SelectTaskById(input.TaskId)

	form := S12Form {
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

