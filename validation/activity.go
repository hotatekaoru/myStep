package validation

import (
	"time"
)

type S11Form struct {
	Date        time.Time `form:"date" validate:"required"`
	TaskId      int       `form:"taskId"`
}

type S12Form struct {
	New         bool    `form:"new"`
	TaskId      int     `form:"taskId"`
	TypeId      int     `form:"typeId" validate:"required,gte=1,lte=3"`
	ContentName string  `form:"contentName" validate:"required"`
	Point       float64 `form:"point" validate:"required,lte=10"`
	Par         int     `form:"workingTime" validate:"required,gte=1,lte=100"`
	UnitId      int     `form:"unitId" validate:"required,gte=1,lte=2"`
}

