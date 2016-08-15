package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Users users = users{}

type users struct{}

func (u *users) U01G01(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})

}
