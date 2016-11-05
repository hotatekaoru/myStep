package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/model"
	"github.com/hotatekaoru/myStep/session"
	"github.com/hotatekaoru/myStep/validation"
	"net/http"
)

type error interface {
	Error() string
}

/* ログイン画面表示処理 */
func S01B01(c *gin.Context) {
	user := session.GetSessionUser(c.Request)
	if (model.User{}) != user {
		session.Destroy(c)
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"userName": "",
	})
}

/* ログイン処理 */
func S01B02(c *gin.Context) {
	input, err := validation.ValidateS01Form(c)
	if err != nil {
		return
	}

	userID, err := model.Auth(input)
	if err != nil {
		model.AuthErr(input.UserName, c)
		return
	}

	session.SaveUserID(c, userID)

	form := model.GetDashBoardInfo(userID)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"form": form,
	})
}

/* Dashboard表示処理 */
func S02B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetDashBoardInfo(user.Id)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"form": form,
	})
}

func J01B01(c *gin.Context) {

	println("Accept JSON Parameter to Login")
	println("userName -> " + c.PostForm("userName"))
	println("password -> " + c.PostForm("password"))

	input, err := validation.ValidateS01FormFromJSON(c)
	if err != nil {
		println("Validate err -> " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			"id": 0,
		})
		return
	}

	userID, err := model.Auth(input)
	if err != nil {
		println("Auth err -> " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
			"id": 0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"err": "",
		"id": userID,
	})
}
