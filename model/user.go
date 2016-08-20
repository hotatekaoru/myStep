package model

import (
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"myStep/constant"
	"myStep/database"
	"question/form"
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

	db.Where(&User{UserName: form.UserName}).Find(&user)
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
/*
func (u *User) BeforeSave() {
	token, err := getUUID(u.UserName)
	if err != nil { panic(err) }
	u.Token = token
}
*/


func getUserID(signature string) (string, error) {
	var uid string
	u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(signature))
	if err == nil {
		uid = u5.String()
	} else {
		uid = ""
	}

	return uid, err
}

func CountUserNameById(userForm *form.UserForm) int {

	var count int
	db.Debug().Model(&User{}).Where("user_name = ?", userForm.UserName).Count(&count)

	return count

}

func CreateUser(userForm *form.UserForm) {

	user := User{
		UserName:  userForm.UserName,
		Password:  PasswordHash(userForm.Password),
	}
	db.Create(&user)

}

func GetUserNameByID(id uint) string {

	var users []User
	db.Debug().Model(&User{}).Where("id = ?", id).Find(&users)

	return users[0].UserName

}

func GetAllUsers() []User {

	var users []User
	db.Debug().Order("id").Find(&users)

	return users

}
