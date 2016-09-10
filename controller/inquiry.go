package controller

import (
	"github.com/gin-gonic/gin"
	"myStep/constant"
	"myStep/model"
	"myStep/session"
	"myStep/validation"
	"net/http"
)

/* アクティビティ一覧照会画面表示処理 */
func (u *users) S21B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS21Form(user.Id)

	c.HTML(http.StatusOK, "activity_table.html", gin.H{
		"form": form,
	})
}

/* アクティビティ一覧照会画面検索処理 */
func (u *users) S21B02(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateS21Form(c)
	form := model.ConvertInputFormToInquiry(input)
	if err != nil {
		model.ReturnS21InputErr(*form, err, c)
		return
	}

	session.SaveS21Form(c, *input)
	form = model.GetS21Search(form)

	c.HTML(http.StatusOK, "activity_table.html", gin.H{
		"form": form,
	})
}

/* アクティビティ削除処理 */
func (u *users) S21B04(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	activityId := validation.ValidateS11URLQueryAcitivityId(c)

	model.DeleteActivityById(activityId)

	input := session.GetSessionInquiry(c.Request)

	form := model.ConvertInputFormToInquiry(&input)
	form = model.GetS21Search(form)

	c.HTML(http.StatusOK, "activity_table.html", gin.H{
		"form": form,
		"info": constant.MSG_COMPLETE_ACTIVITY_DELETE,
	})
}
