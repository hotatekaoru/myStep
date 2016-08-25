package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myStep/model"
)


/* タスクテーブル画面表示処理 */
func (u *users) S41B01(c *gin.Context, ) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	taskList := model.SelectAllTask()
	form := model.ConvTaskListToS41Form(taskList)

	c.HTML(http.StatusOK, "task_table.html", gin.H{
		"form"	: form,
	})
}

/* タスク登録画面表示処理 */
func (u *users) S41B02(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	form := model.GetS42FormRegister()

	c.HTML(http.StatusOK, "task_register.html", gin.H{
		"form": *form,
	})
}

/* タスク登録処理*/
func (u *users) S42P01(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	u.S41B01(c)
}