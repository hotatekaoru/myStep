package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myStep/model"
)

/* タスクテーブル画面表示処理 */
func (u *users) S11B01(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS11FormRegister()

	c.HTML(http.StatusOK, "activity_register.html", gin.H{
		"form": form,
	})
}

func (u *users) S11B02(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS11FormRegister()

	c.HTML(http.StatusOK, "activity_confirm.html", gin.H{
		"form": form,
	})
}
