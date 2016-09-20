package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/model"
	"github.com/hotatekaoru/myStep/session"
	"github.com/hotatekaoru/myStep/validation"
	"net/http"
)

/* アクティビティ登録画面1表示処理 */
func (u *users) S11B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	session.DeinitActivity(c)
	typeId := validation.ValidateS11URLQuery(c)

	form := model.GetS11FormRegister(typeId, user.Id)

	c.HTML(http.StatusOK, "activity_register1.html", gin.H{
		"form": form,
	})
}

/* アクティビティ登録画面2表示処理 */
func (u *users) S11B02(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateS11Form(c)
	if err != nil {
		model.ReturnS11InputErr(input, err, c)
		return
	}

	validation.SetTaskId(input)
	session.SaveS11Form(c, *input)

	form := model.GetS12FormRegister(input)

	c.HTML(http.StatusOK, "activity_register2.html", gin.H{
		"form": form,
	})
}

/* アクティビティ登録処理 */
func (u *users) S12B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateS12Form(c)
	s11 := session.GetSessionActivity(c.Request)
	if err != nil {
		model.ReturnS12InputErr(&s11, input, err, c)
		return
	}

	if s11.New {
		model.CreateActivity(&s11, input)
	} else {
		model.UpdateActivity(&s11, input)
	}
	session.DeinitActivity(c)

	form := model.GetDashBoardInfo(user.Id)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"form": form,
	})
}

/* アクティビティ登録画面1表示（戻る）処理 */
func (u *users) S12B02(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	s := session.GetSessionActivity(c.Request)
	form := model.GetS11FormBySession(s)

	c.HTML(http.StatusOK, "activity_register1.html", gin.H{
		"form": form,
	})
}

/* アクティビティ登録画面1表示処理（更新） */
func (u *users) S21B03(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	session.DeinitActivity(c)
	activityId := validation.ValidateS11URLQueryAcitivityId(c)

	form := model.GetS11FormRegisterByActivityId(activityId)

	c.HTML(http.StatusOK, "activity_register1.html", gin.H{
		"form": form,
	})
}
