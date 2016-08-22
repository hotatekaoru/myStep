package model

import (
	"golang.org/x/crypto/bcrypt"
	"myStep/constant"
	"myStep/database"
	"myStep/validation"
	"errors"
	"time"
)

// Gormのデフォルトでは、IDをunit型にしているが、
// 変換が面倒、かつ、intの最大値2147483647を超過する予定はないので、intで実装する
type GormModel struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	GormModel
	UserName  string `gorm:"not null;unique;type:varchar(100)"`
	Password  string `gorm:"not null;type:varchar(200)"`
}

var db = database.GetDB()

func Auth(form *validation.S01Form) (int, error) {
	user := User{}

	db.Debug().Model(&User{}).Where(&User{UserName: form.UserName}).Find(&user)
	if (User{}) == user { return 0, errors.New(constant.MSG_ENABLE_LOGIN) }

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil { err = errors.New(constant.MSG_ENABLE_LOGIN) }
	return user.ID, err
}

func PasswordHash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil { panic(err) }
	return string(hashed)
}