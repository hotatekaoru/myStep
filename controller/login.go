package controller

import (
	"github.com/gin-gonic/gin"
	"myStep/model"
	"myStep/session"
	"myStep/validation"
	"net/http"
)

var Users users = users{}

type users struct{}

/* ログイン画面表示処理 */
func (u *users) S01B01(c *gin.Context) {
	user := session.GetSessionUser(c.Request)
	if (model.User{}) != user {
		session.Destroy(c)
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"userName": "",
	})
}

/* ログイン処理 */
func (u *users) S01B02(c *gin.Context) {
	form, err := validation.ValidateUser(c)
	if err != nil {
		return
	}

	userID, err := model.Auth(form)
	if err != nil {
		model.AuthErr(form.UserName, c)
		return
	}

	session.SaveUserID(c, userID)

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* Dashboard表示処理 */
func (u *users) S02B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{})
}

/* ログアウト処理 */
func (u *users) S02B02(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"userName": "",
	})
}
