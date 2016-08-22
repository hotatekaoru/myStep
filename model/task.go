package model

import "strconv"

type Task struct {
	GormModel
	TypeId      int `gorm:"not null:numeric(1,0)"`
	ContentId   int `gorm:"not null:numeric(2,0)"`
	ContentName string `gorm:"not null:varchar(100)"`
	Point       float64 `gorm:"not null:numeric(5,1)"`
	Par         int `gorm:"not null:numeric(2,0)"`
	PointUnit   int `gorm:"not null:numeric(1,0)"`
}

type S41Form struct {
	typeName string
	contentName string
	point string
}

var typeMap = map[int]string{1:"Coding", 2:"Training", 3:"Housework"}
var pointUnitMap = map[int]string{1:"回", 2:"分"}

func SelectAllTask() *[]Task{
	task := []Task{}
	db.Debug().Model(&Task{}).Order("type_id, content_id").Find(&task)
	return &task
}

func ConvTaskListToS41Form(taskList *[]Task) *[]S41Form {
	var form []S41Form
	for _, task := range *taskList {
		form = append(form, convOneTask(task))
	}
	return &form
}

func convOneTask(task Task) S41Form {
	form := S41Form{}

	s, _ := typeMap[task.TypeId]
	form.typeName = s

	form.contentName = task.ContentName

	s, _ = pointUnitMap[task.Par]
	form.point = strconv.FormatFloat(task.Point, 'e', 2, 64) + "pt / " +
		strconv.Itoa(task.Par) + s

	return form
}


