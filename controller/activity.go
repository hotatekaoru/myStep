package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myStep/model"
	"myStep/validation"
)

/* タスク登録画面1表示処理 */
func (u *users) S11B01(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	typeId := validation.ValidateS11URLQuery(c)

	form := model.GetS11FormRegister(typeId)

	c.HTML(http.StatusOK, "activity_register1.html", gin.H{
		"form": form,
	})
}

/* タスク登録画面2表示処理 */
func (u *users) S11B02(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateS11Form(c)
	if err != nil {
		model.ReturnS11InputErr(input, err, c)
		return
	}

	form := model.GetS12FormRegister()

	c.HTML(http.StatusOK, "activity_register2.html", gin.H{
		"form": form,
	})
}

/* タスク登録処理 */
func (u *users) S12B01(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	_,_ = validation.ValidateS12Form(c)
	form := model.GetS12FormRegister()

	c.HTML(http.StatusOK, "activity_register2.html", gin.H{
		"form": form,
	})
}
