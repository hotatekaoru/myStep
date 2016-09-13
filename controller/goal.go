package controller

import (
	"github.com/gin-gonic/gin"
	"myStep/model"
	"myStep/session"
	"net/http"
)

/* アクティビティ一覧照会画面表示処理 */
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