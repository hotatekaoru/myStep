package model


type Goal struct {
	GormModel
	UserId     int     `gorm:"not null"`
	Month      string  `gorm:"not null"`
	Coding     int     `gorm:"not null"`
	Training   int     `gorm:"not null"`
	HouseWork  int     `gorm:"not null"`
}

type S31Form struct {
	goal []Goal
}

/* DB操作 */

/* form（外部出力）操作 */
func GetS31Form() {
	
}