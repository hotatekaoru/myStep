package main

import (
	"github.com/gin-gonic/gin"
	"myStep/controller"
	"myStep/database"
	"myStep/model"
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
	router.GET("/login", controller.Users.S01B01)
	router.POST("/index", controller.Users.S01B02)

	/* S02_Dashboard処理 */
	router.GET("/", controller.Users.S02B01)
	router.GET("/index", controller.Users.S02B01)

	/* S11_アクティビティ登録画面1処理 */
	router.GET("/activity/register/typeId=:typeId", controller.Users.S11B01)
	router.POST("/activity/confirm", controller.Users.S11B02)

	/* S12_アクティビティ登録画面2処理 */
	router.POST("/activity/complete", controller.Users.S12B01)
	router.GET("/activity/register", controller.Users.S12B02)

	/* S21_アクティビティ一覧照会画面処理 */
	router.GET("/activity/inquiry", controller.Users.S21B01)

	/* S41_タスクテーブル画面処理 */
	router.GET("/task_table", controller.Users.S41B01)
	router.GET("/task/register", controller.Users.S41B02)
	router.POST("/task/update", controller.Users.S41B03)
	router.POST("/task/delete", controller.Users.S41B04)

	/* S42_タスクテーブル画面処理 */
	router.POST("/task/register", controller.Users.S42B01)
	http.ListenAndServe(":"+port(), router)
}

/* DBの自動生成 */
func migrate() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.Activity{})
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}
