package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func YzmChnageapi(c *gin.Context) {
	yzm:=[]string{"1.png","2.jpg"}
	temp:="assets/img/yzm/"+yzm[rand.Intn(2)]
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"yzm": temp,
	})
}
