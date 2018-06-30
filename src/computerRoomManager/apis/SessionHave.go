package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func isHaveSessions(c *gin.Context)(id string) {
	session := sessions.Default(c)
	v := session.Get("id")
	if(v==nil){
		c.Redirect(http.StatusMovedPermanently,"/")
		return ""
	}
	if(session.Get("flag").(string)=="m"){
         if(v.(string)[0]!='m'){
			 c.Redirect(http.StatusMovedPermanently,"/")
			 return ""
		 }
	}
	return v.(string)
}