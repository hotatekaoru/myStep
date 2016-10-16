package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/model"
	"github.com/hotatekaoru/myStep/session"
	"net/http"
)

/* 体重登録画面表示処理 */
func S51B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS51FormRegister()

	c.HTML(http.StatusOK, ".html", gin.H{
		"form": form,
	})
}

/* 体重一覧画面表示処理 */
func S52B01(c *gin.Context) {
	user := session.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	c.HTML(http.StatusOK, ".html", gin.H{
	})
}

