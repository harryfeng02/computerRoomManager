package apis

import (
	db "computerRoomManager/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Undo(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	applyno := c.PostForm("applyno")
	var temp=""
	error:=db.MyDB.QueryRow("select tno from teach where tid in (select tid from apply where applyno=?)",applyno).Scan(&temp)
	if(temp!=tno){
		fmt.Println(temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		fmt.Println("vc")
		//_,error:=db.MyDB.Exec("delete from apply where applyno=?",applyno)
		//if error == nil {
		//	c.JSON(http.StatusOK, "ok")
		//}else {
		//	c.JSON(http.StatusOK, error)
		//}
	}

}
