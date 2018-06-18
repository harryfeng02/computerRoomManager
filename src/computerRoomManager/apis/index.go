package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"fmt"
)

func PreLoginApi(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func IndexApi(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HomeInitApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func DetailApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "detail.html", gin.H{})
}
func CheckApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "check.html", gin.H{})
}
func MyClassApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "myclass.html", gin.H{})
}
func MangerHomeInitApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "manage.html", gin.H{})
}
func MangerCheckApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "checkApply.html", gin.H{})
}
func MangerUpdateApi(c *gin.Context) {
	isHaveSessions(c)
	c.HTML(http.StatusOK, "checkSoftware.html", gin.H{})
}

func LogoutApi(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("id")
	if(v!=nil){
		session.Clear()
		session.Save()
		fmt.Println("logout")
	}
	c.JSON(http.StatusOK,"ok")
}




