package model

import (
	"github.com/hotatekaoru/myStep/constant"
	"time"
)

type WeightMemory struct {
	GormModel
	Date        time.Time `gorm:"not null"`
	Weight      float64   `gorm:"not null"`
	FatPar      float64   `gorm:"not null"`
	FatKg       float64   `gorm:"not null"`
	Comment     string    `gorm:"varchar(200)"`
}

/* DB操作 */
type S51Form struct {
	New     bool
	Date    string
	Weight  float64
	FatPar  float64
	Comment string
}

/* form（外部出力）操作 */
func GetS51FormRegister() *S51Form {
	form := S51Form {
		New:    true,
		Date:   time.Now().Format(constant.DATE_FORMAT_YYYYMMDD),
	}

	return &form
}