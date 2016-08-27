package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myStep/model"
	"myStep/validation"
	"myStep/constant"
)

/* タスクテーブル画面表示処理 */
func (u *users) S41B01(c *gin.Context) {

	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	taskList := model.SelectAllTask()
	form := model.ConvTaskListToS41Form(taskList)

	c.HTML(http.StatusOK, "task_table.html", gin.H{
		"form": form,
	})
}

/* タスク登録画面表示処理（新規登録） */
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

/* タスク登録画面表示処理（更新） */
func (u *users) S41B03(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	taskId, err := validation.ValidateTaskId(c)
	if err != nil {
		model.ForceLogOut(c)
		return
	}

	task := model.SelectTaskById(taskId)
	form := model.GetS42FormUpdate(task)

	c.HTML(http.StatusOK, "task_register.html", gin.H{
		"form": *form,
	})
}

/* タスク削除処理 */
func (u *users) S41B04(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	taskId, err := validation.ValidateTaskId(c)
	if err != nil {
		model.ForceLogOut(c)
		return
	}

	model.DeleteTaskById(taskId)
	msg = constant.MSG_COMPLETE_TASK_DELETE

	taskList := model.SelectAllTask()
	form := model.ConvTaskListToS41Form(taskList)

	c.HTML(http.StatusOK, "task_table.html", gin.H{
		"form": form,
		"info": msg,
	})
}

/* タスク登録・更新処理 */
func (u *users) S42B01(c *gin.Context) {
	user := model.IsLogin(c)
	if (model.User{}) == user {
		return
	}

	input, err := validation.ValidateTask(c)
	if err != nil {
		model.ReturnS42InputErr(input, err, c)
		return
	}

	var msg string
	if input.New {
		model.CreateTask(input)
		msg = constant.MSG_COMPLETE_TASK_REGISTER
	} else {
		model.UpdateTask(input)
		msg = constant.MSG_COMPLETE_TASK_UPDATE
	}

	taskList := model.SelectAllTask()
	form := model.ConvTaskListToS41Form(taskList)

	c.HTML(http.StatusOK, "task_table.html", gin.H{
		"form": form,
		"info": msg,
	})
}