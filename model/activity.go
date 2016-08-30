package model

import "time"

// 一部タスクテーブルと同じ構成だが、Taskテーブルは変更を許容しているため、
// Activity登録時点のTask情報をActivityテーブルに持つ
type Activity struct {
	GormModel
	UserId      int `gorm:"not null:numeric(1,0)"`
	// ジャンル 1:Coding, 2:Training, 3:Housework
	TypeId      int `gorm:"not null:numeric(1,0)"`
	ContentId   int `gorm:"not null:numeric(3,0)"`
	ContentName string `gorm:"not null:varchar(100)"`
	Point       float64 `gorm:"not null:numeric(5,1)"`
	// 単位（数）
	Par         int `gorm:"not null:numeric(2,0)"`
	// 1:回, 2:分
	UnitId      int `gorm:"not null:numeric(1,0)"`
	Comment     string `gorm:"varchar(300)"`
}

type contents struct {
	ContentId   int
	ContentName string
}

type S11Form struct {
	New              bool
	TypeId           int
	Point            float64
	Par              int
	UnitId           int
	CodingList       []contents
	TrainingList     []contents
	HouseworkList    []contents
}

type S12Form struct {
	New              bool
	Date             time.Time
	TaskName         string
	ContentName      string
	Point            float64
	WorkingTime      int
	UnitId           int
}

/* DB操作 */

/* form（外部出力）操作 */
func GetS11FormRegister() *S11Form{

	form := S11Form {
		New:         true,
		TypeId:      1,
		Point:       1.0,
		Par:         1,
		UnitId:      1,
	}
	tasks := SelectAllTask()
	setTaskToS11Form(&form, tasks)

	return &form
}

func setTaskToS11Form(form *S11Form, tasks *[]Task) {
	for _, v := range *tasks {
		content := contents {
			ContentId:   v.ContentId,
			ContentName: v.ContentName,
		}
		switch v.TypeId {
		case 1:
			form.CodingList = append(form.CodingList, content)
		case 2:
			form.TrainingList = append(form.TrainingList, content)
		case 3:
			form.HouseworkList = append(form.HouseworkList, content)
		}
	}
}

func GetS12FormRegister() *S12Form{

	form := S12Form {
		New:         true,
		Date:        time.Now(),
		TaskName:    "Coding",
		ContentName: "tidy",
		Point:       1.0,
		WorkingTime: 1,
		UnitId:      1,
	}
	return &form
}

