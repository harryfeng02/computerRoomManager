package apis

import (
	"github.com/gin-gonic/gin"
	db "computerRoomManager/database"
	"net/http"
)

func Undo(c *gin.Context) {
	applyno := c.PostForm("applyno")
	_,error:=db.MyDB.Exec("delete from apply where applyno=?",applyno)
	if error == nil {
		c.JSON(http.StatusOK, "ok")
	}else {
		c.JSON(http.StatusOK, error)
	}
}
