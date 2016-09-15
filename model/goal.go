package model

import (
	"myStep/validation"
	"time"
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
	Total     int
	New       bool
}

/* DB操作 */
func selectAllGoals(userId int) []Goal {
	goal := []Goal{}
	db.Debug().Model(&Goal{}).Where("user_id = ?", userId).Order("Month").Find(&goal)
	return goal
}

func createGoal(input *validation.S32Form) {
	goal := Goal{
		UserId:    input.UserId,
		Month:     input.Year + input.Month,
		Coding:    input.Coding,
		Training:  input.Training,
		HouseWork: input.HouseWork,
		Total:     input.Coding + input.Training + input.HouseWork,
	}

	db.Debug().Create(&goal)
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
		Total:     90,
		New:       true,
	}

	return form
}

func getYearList() []int {
	year := time.Now().Year()
	return []int{year, year + 1, year + 2}
}
