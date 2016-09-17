package controller

import (
	"github.com/gin-gonic/gin"
	"myStep/model"
	"myStep/session"
	"net/http"
	"myStep/validation"
)

/* 目標一覧画面表示処理 */
func (u *users) S31B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS31Form(user.Id)

	c.HTML(http.StatusOK, "goal_table.html", gin.H{
		"form": form,
	})
}

/* 目標登録画面表示処理 */
func (u *users) S31B02(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS32Form()

	c.HTML(http.StatusOK, "goal_register.html", gin.H{
		"form": form,
	})
}

/* 目標登録処理 */
func (u *users) S32B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateS32Form(c)
	if err != nil {
		model.ReturnS32InputErr(input, err, c)
		return
	}
	input.UserId = user.Id

	model.RegisterGoal(input)

	form := model.GetS31Form(user.Id)

	c.HTML(http.StatusOK, "goal_table.html", gin.H{
		"form": form,
	})
}
