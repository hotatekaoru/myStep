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
	gob.Register(validation.S11Form{})

	var err error
	max := maxAge()
	store, err = NewRediStoreWithPool(database.GetRedisPool(), []byte("secret-key"))
	if err != nil {
		panic(err)
	}

	store.SetMaxAge(max)
}

func GetSessionActivity(req *http.Request) validation.S11Form {
	activity, ok := (GetSession(req).Values[constant.SESSION_ACTIVITY]).(validation.S11Form)
	if !ok {
		return validation.S11Form{}
	}
	return activity
}

func SaveS11Form(c *gin.Context, form validation.S11Form) {
	s := GetSession(c.Request)
	s.Values[constant.SESSION_ACTIVITY] = form
	Save(c.Request, c.Writer)
}

func DeinitActivity(c *gin.Context) {
	SaveS11Form(c, validation.S11Form{})
}
