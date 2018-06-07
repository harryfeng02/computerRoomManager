package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
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


