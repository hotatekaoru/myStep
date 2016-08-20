package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
	"myStep/controller"
	"net/http"
	"os"
	"unicode"
	"myStep/database"
	"myStep/model"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
	msgValidationFailed = func(e *validator.FieldError) string {
		switch e.Tag {
		case "required":
			return toSnakeCase(e.Field) + ": required"
		case "max":
			return toSnakeCase(e.Field) + ": too_long"
		default:
			return e.Error()
		}
	}
)

func main() {

	// DBの自動生成
	migrate()

	router := gin.Default()
	router.Static("/assets", "./assets/")
	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	/* B01_ログイン画面処理 */
	router.GET("/login", controller.Users.S01B01)
	router.POST("/index", controller.Users.S01B02)

	/* B02_Dashboard処理 */
	router.GET("/", controller.Users.S02B01)
	router.GET("/index", controller.Users.S02B01)

	http.ListenAndServe(":"+port(), router)
}

/* DBの自動生成 */
func migrate() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
}

// https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}
