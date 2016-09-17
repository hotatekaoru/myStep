package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"github.com/hotatekaoru/myStep/constant"
	"github.com/hotatekaoru/myStep/validation"
	"net/http"
	"strconv"
)

type Task struct {
	GormModel
	// ジャンル 1:Coding, 2:Training, 3:Housework
	TypeId  int     `gorm:"not null"`
	Content string  `gorm:"not null"`
	Point   float64 `gorm:"not null"`
	// 1:1回, 2:10分
	UnitId int `gorm:"not null"`
}

type S41Form struct {
	TaskId   int
	TypeName string
	Content  string
	PointStr string
}

type S42Form struct {
	New     bool
	TaskId  int
	TypeId  int
	Content string
	Point   float64
	UnitId  int
}

var typeMap = map[int]string{1: "Coding", 2: "Training", 3: "Housework"}
var pointUnitMap = map[int]string{1: "1回", 2: "10分"}

/* DB操作 */
func SelectAllTask() *[]Task {
	task := []Task{}
	db.Debug().Model(&Task{}).Order("type_id, id").Find(&task)
	return &task
}

func CreateTask(input *validation.S42Form) {
	task := Task{
		TypeId:  input.TypeId,
		Content: input.Content,
		Point:   input.Point,
		UnitId:  input.UnitId,
	}

	db.Debug().Create(&task)
}

func SelectTaskById(id int) *Task {
	task := Task{}
	db.Debug().Model(&Task{}).First(&task, id)
	return &task
}

func UpdateTask(input *validation.S42Form) {
	task := SelectTaskById(input.TaskId)
	if (Task{}) == *task {
		return
	}

	db.Debug().Model(&task).Updates(Task{
		TypeId:  input.TypeId,
		Content: input.Content,
		Point:   input.Point,
		UnitId:  input.UnitId,
	})
}

func DeleteTaskById(id int) {
	db.Debug().Where("id = ?", id).Delete(&Task{})
}

/* form（外部出力）操作 */
func ConvTaskListToS41Form(taskList *[]Task) *[]S41Form {
	var form []S41Form
	for _, task := range *taskList {
		form = append(form, convOneTaskToS41(task))
	}
	return &form
}

func convOneTaskToS41(task Task) S41Form {
	form := S41Form{}

	form.TaskId = task.Id
	s, _ := typeMap[task.TypeId]
	form.TypeName = s

	form.Content = task.Content

	s, _ = pointUnitMap[task.UnitId]
	form.PointStr = strconv.FormatFloat(task.Point, 'f', -2, 64) + "pt / " + s

	return form
}

func GetS42FormRegister() *S42Form {
	form := S42Form{
		New:     true,
		TypeId:  1,
		Content: "",
		Point:   1.0,
		UnitId:  1,
	}
	return &form
}

func ReturnS42InputErr(input *validation.S42Form, errs error, c *gin.Context) {
	form := S42Form{
		New:     input.New,
		TypeId:  input.TypeId,
		Content: input.Content,
		Point:   input.Point,
		UnitId:  input.UnitId,
	}

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field()+constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "task_register.html", gin.H{
		"errorList": err,
		"form":      form,
	})
}

func GetS42FormUpdate(task *Task) *S42Form {
	form := S42Form{
		New:     false,
		TaskId:  task.Id,
		TypeId:  task.TypeId,
		Content: task.Content,
		Point:   task.Point,
		UnitId:  task.UnitId,
	}
	return &form
}
