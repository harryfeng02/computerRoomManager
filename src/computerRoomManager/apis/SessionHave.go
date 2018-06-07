package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"fmt"
	"net/http"
)

func isHaveSessions(c *gin.Context)(id string) {
	session := sessions.Default(c)
	v := session.Get("id")
	fmt.Println(v,"test1")
	if(v==nil){
		c.Redirect(http.StatusMovedPermanently,"/")
		return ""
	}
	return "000000"
}