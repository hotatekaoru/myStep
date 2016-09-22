package session

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/hotatekaoru/myStep/constant"
	"github.com/hotatekaoru/myStep/database"
	"github.com/hotatekaoru/myStep/validation"
	. "gopkg.in/boj/redistore.v1"
	"net/http"
)

func init() {
	gob.Register(validation.S21Form{})

	var err error
	max := maxAge()
	store, err = NewRediStoreWithPool(database.GetRedisPool(), []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	store.SetMaxAge(max)
}

func GetSessionInquiry(req *http.Request) validation.S21Form {
	activity, ok := (GetSession(req).Values[constant.SESSION_INQUIRY]).(validation.S21Form)
	if !ok {
		return validation.S21Form{}
	}
	return activity
}

func SaveS21Form(c *gin.Context, form validation.S21Form) {
	s := GetSession(c.Request)
	s.Values[constant.SESSION_INQUIRY] = form
	Save(c.Request, c.Writer)
}
