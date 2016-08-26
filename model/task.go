package model

import (
	"strconv"
	"myStep/validation"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"myStep/constant"
	"errors"
)

type Task struct {
	GormModel
	// ジャンル 1:Coding, 2:Training, 3:Housework
	TypeId      int `gorm:"not null:numeric(1,0)"`
	ContentId   int `gorm:"not null:numeric(3,0)"`
	ContentName string `gorm:"not null:varchar(100)"`
	Point       float64 `gorm:"not null:numeric(5,1)"`
	// 単位（数）
	Par         int `gorm:"not null:numeric(2,0)"`
	// 1:回, 2:分
	UnitId   int `gorm:"not null:numeric(1,0)"`
	// true:存在, false:削除
	// GormModelのdeleteTimeでは、削除した場合に集計からも削除される。
	// UsedFlgは、削除時点から該当タスクの登録を行えないようにするが、
	// 過去に遡っての集計は可能にするためのフラグ
	UsedFlg     bool `gorm:"not null;"`
}

type S41Form struct {
	TaskId int
	TypeName string
	ContentName string
	PointStr string
}

type S42Form struct {
	New bool
	TaskId int
	TypeId int
	ContentName string
	Point float64
	Par int
	UnitId int
}

var typeMap = map[int]string{1:"Coding", 2:"Training", 3:"Housework"}
var pointUnitMap = map[int]string{1:"回", 2:"分"}

/* DB操作 */
func SelectAllTask() *[]Task{
	task := []Task{}
	db.Debug().Model(&Task{}).Order("type_id, content_id").Find(&task)
	return &task
}

func CreateTask(input *validation.S42Form) {
	task := Task{
		TypeId:      input.TypeId,
		ContentId:   1,
		ContentName: input.ContentName,
		Point:       input.Point,
		Par:         input.Par,
		UnitId:      input.UnitId,
		UsedFlg:     true,
	}

	db.Debug().Create(&task)

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

	form.ContentName = task.ContentName

	s, _ = pointUnitMap[task.Par]
	form.PointStr = strconv.FormatFloat(task.Point, 'e', 2, 64) + "pt / " +
		strconv.Itoa(task.Par) + s

	return form
}

func GetS42FormRegister() *S42Form{
	form := S42Form {
		New: true,
		TypeId: 1,
		ContentName: "",
		Point: 1.0,
		Par: 1,
		UnitId: 1,
	}
	return &form
}

func ReturnS42InputErr(input *validation.S42Form, errs error,c *gin.Context) {
	form := S42Form {
		New: input.New,
		TypeId: input.TypeId,
		ContentName: input.ContentName,
		Point: input.Point,
		Par: input.Par,
		UnitId: input.UnitId,
	}

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field + constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "task_register.html", gin.H{
		"errorList": err,
		"form": form,
	})
}


