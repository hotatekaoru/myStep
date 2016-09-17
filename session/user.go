package session

import (
	"encoding/gob"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	. "gopkg.in/boj/redistore.v1"
	"github.com/hotatekaoru/myStep/constant"
	"github.com/hotatekaoru/myStep/database"
	"github.com/hotatekaoru/myStep/model"
	"net/http"
	"os"
	"strconv"
)

var (
	store *RediStore
)

func init() {
	gob.Register(model.User{})

	var err error
	max := maxAge()
	store, err = NewRediStoreWithPool(database.GetRedisPool(), []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	store.SetMaxAge(max)
}

func maxAge() int {
	env := os.Getenv("SESSION_MAX_AGE")
	if env == "" {
		return constant.SESSION_MAX_AGE
	}
	max, _ := strconv.Atoi(env)

	return max
}

func GetSession(req *http.Request) *sessions.Session {

	session, err := store.Get(req, constant.SESSION_KEY)
	if err != nil {
		panic(err)
	}
	return session
}

func GetSessionUser(req *http.Request) model.User {
	userID := (GetSession(req).Values[constant.SESSION_USER_ID])
	if userID == nil {
		return model.User{}
	}

	var user model.User
	db := database.GetDB()
	db.First(&user, userID)
	return user
}

func Save(req *http.Request, res http.ResponseWriter) {
	if err := sessions.Save(req, res); err != nil {
		panic(err)
	}
}

func SaveUserID(c *gin.Context, userID int) {
	s := GetSession(c.Request)
	s.Values[constant.SESSION_USER_ID] = userID
	Save(c.Request, c.Writer)
}

func IsLogin(c *gin.Context) model.User {
	user := GetSessionUser(c.Request)
	if (model.User{}) == user {

		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"userName": "",
			"error": []error{
				errors.New(constant.MSG_ENABLE_GET_USER_DATA),
			},
		})
	}
	return user
}

func Destroy(c *gin.Context) {
	session := GetSession(c.Request)
	session.Options.MaxAge = -1
	Save(c.Request, c.Writer)
}
