package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/controller"
	"github.com/hotatekaoru/myStep/database"
	"github.com/hotatekaoru/myStep/model"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	// DBの自動生成
	migrate()

	router := gin.Default()
	router.Static("/assets", "./assets/")
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	/* S01_ログイン画面処理 */
	router.GET("/login", controller.S01B01)
	router.POST("/index", controller.S01B02)

	/* S02_Dashboard処理 */
	router.GET("/", controller.S02B01)
	router.GET("/index", controller.S02B01)

	/* S11_アクティビティ登録画面1処理 */
	router.GET("/activity/register/typeId=:typeId", controller.S11B01)
	router.POST("/activity/confirm", controller.S11B02)

	/* S12_アクティビティ登録画面2処理 */
	router.POST("/activity/complete", controller.S12B01)
	router.GET("/activity/register", controller.S12B02)

	/* S21_アクティビティ一覧照会画面処理 */
	router.GET("/activity/inquiry", controller.S21B01)
	router.POST("/activity/inquiry", controller.S21B02)
	router.GET("/activity/register/activityId=:activityId", controller.S21B03)
	router.GET("/activity/delete/activityId=:activityId", controller.S21B04)
	router.GET("/activity/inquiry/typeId=:typeId", controller.S21B05)

	/* S31_目標一覧画面処理 */
	router.GET("/goal/list", controller.S31B01)
	router.GET("/goal/register", controller.S31B02)
	router.GET("/goal/register/month=:month", controller.S31B03)

	/* S32_目標登録画面処理 */
	router.POST("/goal/register", controller.S32B01)

	/* S41_タスクテーブル画面処理 */
	router.GET("/task_table", controller.S41B01)
	router.GET("/task/register", controller.S41B02)
	router.POST("/task/update", controller.S41B03)
	router.POST("/task/delete", controller.S41B04)

	/* S42_タスクテーブル画面処理 */
	router.POST("/task/register", controller.S42B01)

	/* S51_体重登録画面処理 */
	router.GET("/weight/register", controller.S51B01)

	/* S52_体重一覧画面処理 */
	router.GET("/weight/inquiry", controller.S52B01)

	/* J01_ユーザー確認処理 */
	router.POST("/JSON/login", controller.J01B01)

	http.ListenAndServe(":"+port(), router)
}

/* DBの自動生成 */
func migrate() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.Activity{})
	db.AutoMigrate(&model.Goal{})
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}
