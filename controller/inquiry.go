package controller

import (
	"github.com/gin-gonic/gin"
	"myStep/model"
	"myStep/session"
	"net/http"
)

/* アクティビティ一覧照会画面表示処理 */
func (u *users) S21B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS21Form()

	c.HTML(http.StatusOK, "activity_table.html", gin.H{
		"form": form,
	})
}