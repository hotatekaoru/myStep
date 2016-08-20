package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myStep/validation"
	"myStep/model"
)

var Users users = users{}

type users struct{}


/* ログイン画面表示処理 */
func (u *users) S01B01(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"userName"	: "",
	})
}

/* ログイン処理 */
func (u *users) S01B02(c *gin.Context) {
	form, err := validation.ValidateUser(c)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"userName"	: c.PostForm("userName"),
			"error"		: []error{
				err,
			},
		})
		return
	}

	userID, err := model.Auth(form)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"userName"	: form.UserName,
			"error"		: []error{
				err,
			},
		})
		return
	}

	model.SaveUserByID(c, userID)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"userName"	: form.UserName,
	})
}

/* Dashboard表示処理 */
func (u *users) S02B01(c *gin.Context) {

	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"userName"	: "",
	})
}


