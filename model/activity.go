package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/constant"
	"github.com/hotatekaoru/myStep/validation"
	"gopkg.in/go-playground/validator.v9"
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

type rate struct {
	Coding    int
	Training  int
	HouseWork int
	Total     int
}

type allTask struct {
	Coding    float64
	Training  float64
	HouseWork float64
	Total     float64
}

type DashForm struct {
	Rate rate
	Now  allTask
	Goal Goal
}

type S11Form struct {
	New           bool
	ActivityId    int
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

type s21Inquiry struct {
	UserList []bool
	DateFrom string
	DateEnd  string
	TypeList []bool
}

type s21Activity struct {
	ActivityId  int
	UserName    string
	Date        string
	TypeName    string
	Content     string
	UnitId      int
	WorkingTime int
	Point       float64
}

type S21Form struct {
	Inquiry  s21Inquiry
	Activity []s21Activity
}

var userMap = map[int]string{1: "Kaoru", 2: "Yuri"}

/* DB操作 */
func selectActivityById(id int) *Activity {
	activity := Activity{}
	db.Debug().Model(&Activity{}).Where("id = ?", id).Find(&activity)
	return &activity
}

func selectActivity(inq *s21Inquiry) *[]Activity {
	activity := []Activity{}
	from, _ := time.Parse(constant.DATE_FORMAT_YYYYMMDD, inq.DateFrom)
	var end time.Time
	if inq.DateEnd == "" {
		end = time.Now()
	} else {
		end, _ = time.Parse(constant.DATE_FORMAT_YYYYMMDD, inq.DateEnd)
	}

	db.Debug().Model(&Activity{}).
		Where("user_id in (?) and (date BETWEEN ? AND ?) and type_id in (?)",
			convertBoolToNumList(inq.UserList), from, end, convertBoolToNumList(inq.TypeList)).
		Order("id").Find(&activity)
	return &activity
}

func selectActivityThisMonth(userId int) *[]Activity {
	YYYYMM := getYYYYMMStr(time.Now().Year(), int(time.Now().Month()))
	from, _ := time.Parse("20060102", YYYYMM+"01")
	end := from.AddDate(0, 1, 0)
	activity := []Activity{}

	db.Debug().Model(&Activity{}).
		Where("user_id = ? and (date BETWEEN ? AND ?)", userId, from, end).Find(&activity)
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

func UpdateActivity(s11 *validation.S11Form, s12 *validation.S12Form) {
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

	db.Debug().Model(&Activity{}).Where("id = ?", s11.ActivityId).Update(&activity)
}

func DeleteActivityById(id int) {
	db.Debug().Where("id = ?", id).Delete(&Activity{})
}

/* form（外部出力）操作 */
func GetDashBoardInfo(userId int) *DashForm {
	form := DashForm{}
	acts := selectActivityThisMonth(userId)
	nowPoint := calcActivityByTask(acts)
	goal := selectGoalByUserAndMonth(userId, getYYYYMMStr(time.Now().Year(), int(time.Now().Month())))
	form.Goal = goal
	form.Now = *nowPoint
	form.Rate = calcRate(nowPoint, &goal)
	return &form
}

func calcActivityByTask(activity *[]Activity) *allTask {
	result := allTask{}
	for _, v := range *activity {
		switch v.TypeId {
		case 1:
			result.Coding += v.Point
		case 2:
			result.Training += v.Point
		case 3:
			result.HouseWork += v.Point
		}
	}
	result.Total = result.Coding + result.Training + result.HouseWork
	return &result
}

func calcRate(nowPoint *allTask, goal *Goal) rate {
	result := rate{
		Coding:    int(100*nowPoint.Coding) / goal.Coding,
		Training:  int(100*nowPoint.Training) / goal.Training,
		HouseWork: int(100*nowPoint.HouseWork) / goal.HouseWork,
		Total:     int(100*nowPoint.Total) / goal.Total,
	}
	return result
}
func GetS11FormRegister(typeId, userId int) *S11Form {
	form := S11Form{
		New:    true,
		UserId: userId,
		Date:   time.Now().Format(constant.DATE_FORMAT_YYYYMMDD),
		TypeId: typeId,
	}
	tasks := SelectAllTask()
	SetTaskToS11Form(&form, tasks)

	return &form
}

func GetS11FormRegisterByActivityId(activityId int) *S11Form {
	activity := selectActivityById(activityId)
	form := S11Form{
		New:        false,
		ActivityId: activityId,
		UserId:     activity.UserId,
		Date:       activity.Date.Format(constant.DATE_FORMAT_YYYYMMDD),
		TypeId:     activity.TypeId,
		TaskId:     activity.TaskId,
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

func GetS21Form(userId int) *S21Form {
	form := S21Form{}

	form.Inquiry = initInquiry(userId)
	activities := selectActivity(&form.Inquiry)
	form.Activity = convertActivitiesToS21Form(activities)

	return &form
}

func GetS21Search(form *S21Form) *S21Form {
	activities := selectActivity(&form.Inquiry)
	form.Activity = convertActivitiesToS21Form(activities)

	return form
}

func initInquiry(userId int) s21Inquiry {
	today := time.Now()
	lastMonth := today.AddDate(0, -1, 0)
	inquiry := s21Inquiry{
		UserList: returnBoolSliceOneTrue(userId-1, 2),
		DateFrom: lastMonth.Format(constant.DATE_FORMAT_YYYYMMDD),
		DateEnd:  today.Format(constant.DATE_FORMAT_YYYYMMDD),
		TypeList: returnBoolSliceAllTrue(3),
	}
	return inquiry
}

func returnBoolSliceOneTrue(choice, len int) []bool {
	slice := make([]bool, len)
	slice[choice] = true
	return slice
}

func returnBoolSliceAllTrue(len int) []bool {
	slice := make([]bool, len)
	for i := range slice {
		slice[i] = true
	}
	return slice
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

func ConvertInputFormToInquiry(obj *validation.S21Form) *S21Form {
	form := S21Form{}
	inquiry := s21Inquiry{
		UserList: convertNumToBoolList(obj.UserCheck, 2),
		DateFrom: obj.DateFrom,
		DateEnd:  obj.DateEnd,
		TypeList: convertNumToBoolList(obj.TypeCheck, 3),
	}
	form.Inquiry = inquiry
	return &form
}

func ReturnS21InputErr(form S21Form, errs error, c *gin.Context) {

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field()+constant.MSG_WRONG_INPUT))
	}

	c.HTML(http.StatusBadRequest, "activity_table.html", gin.H{
		"errorList": err,
		"form":      form,
	})
}

func convertBoolToNumList(boolList []bool) []int {
	numList := []int{}
	for i, v := range boolList {
		if v {
			numList = append(numList, i+1)
		}
	}
	return numList
}

func convertNumToBoolList(numList []int, len int) []bool {
	boolList := make([]bool, len)
	for _, v := range numList {
		boolList[v-1] = true
	}
	return boolList
}
