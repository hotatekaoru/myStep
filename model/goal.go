package model

import (
	"myStep/validation"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"errors"
	"myStep/constant"
	"strconv"
)

type Goal struct {
	GormModel
	UserId    int    `gorm:"not null"`
	Month     string `gorm:"not null"`
	Coding    int    `gorm:"not null"`
	Training  int    `gorm:"not null"`
	HouseWork int    `gorm:"not null"`
	Total     int    `gorm:"not null"`
}

type S31Form struct {
	Goal []Goal
}

type S32Form struct {
	Year      int
	Month     int
	YearList  []int
	MonthList []int
	Coding    int
	Training  int
	HouseWork int
	New       bool
}

/* DB操作 */
func selectAllGoals(userId int) []Goal {
	goal := []Goal{}
	db.Debug().Model(&Goal{}).Where("user_id = ?", userId).Order("month").Find(&goal)
	return goal
}

func isExistGoalByUserAndMonth(userId int, m string) bool {
	goal := []Goal{}
	db.Debug().Model(&Goal{}).Where("user_id = ? and month = ?", userId, m).Find(&goal)
	return len(goal) > 0
}

func createOrUpdateGoal(input *validation.S32Form, m string) {
	if isExistGoalByUserAndMonth(input.UserId, m) {
		updateGoal(input, m)
	} else {
		createGoal(input, m)
	}
}

func createGoal(input *validation.S32Form, m string) {
	goal := Goal{
		UserId:    input.UserId,
		Month:     m,
		Coding:    input.Coding,
		Training:  input.Training,
		HouseWork: input.HouseWork,
		Total:     input.Coding + input.Training + input.HouseWork,
	}

	db.Debug().Create(&goal)
}

func updateGoal(input *validation.S32Form, m string) {
	goal := Goal{
		Coding:    input.Coding,
		Training:  input.Training,
		HouseWork: input.HouseWork,
		Total:     input.Coding + input.Training + input.HouseWork,
	}

	db.Debug().Model(&Goal{}).Where("user_id = ? and month = ?", input.UserId, m).Update(&goal)
}

func RegisterGoal(input *validation.S32Form) {
	monthList := getMonthListForUpdate(input)
	for _, m := range *monthList {
		createOrUpdateGoal(input, m)
	}
}

func getMonthListForUpdate(input *validation.S32Form) *[]string {
	year := convertStrToInt(input.Year)
	month := convertStrToInt(input.Month)

	if input.Continue != 1 {
		return &[]string{getYYYYMMStr(year, month)}
	}
	monthList := []string{}

	for y := time.Now().Year(); y <= time.Now().Year() + 2; y++ {
		if (year > y) {
			continue
		}
		for m := 1; m <= 12; m++ {
			if (month > m) {
				continue
			}
			month = 1
			monthList = append(monthList, getYYYYMMStr(y, m))
		}
	}
	return &monthList
}

func getYYYYMMStr(y, m int) string {
	s := strconv.Itoa(m)
	if len(s) == 1{
		s = "0" + s
	}
	return strconv.Itoa(y) + s
}

/* form（外部出力）操作 */
func GetS31Form(userId int) S31Form {
	form := S31Form{
		Goal: selectAllGoals(userId),
	}

	return form
}

func GetS32Form() S32Form {
	form := S32Form{
		Year:      time.Now().Year(),
		Month:     int(time.Now().Month()),
		YearList:  getYearList(),
		MonthList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		Coding:    30,
		Training:  30,
		HouseWork: 30,
		New:       true,
	}

	return form
}

func getYearList() []int {
	year := time.Now().Year()
	return []int{year, year + 1, year + 2}
}

func ReturnS32InputErr(input *validation.S32Form, errs error, c *gin.Context) {
	form := S32Form{
		New:        input.New,
		Year:       convertStrToInt(input.Year),
		Month:      convertStrToInt(input.Month),
		YearList:   getYearList(),
		MonthList:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		Coding:     input.Coding,
		Training:   input.Training,
		HouseWork:  input.HouseWork,
	}

	var err []error
	for _, v := range errs.(validator.ValidationErrors) {
		err = append(err, errors.New(v.Field()+constant.MSG_WRONG_INPUT))
	}
	c.HTML(http.StatusBadRequest, "goal_register.html", gin.H{
		"errorList": err,
		"form":      form,
	})
}

func convertStrToInt(s string)int {
	i, _ := strconv.Atoi(s)
	return i
}