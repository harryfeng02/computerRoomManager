package apis

import (
	db "computerRoomManager/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-contrib/sessions"
)

func Undo(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	applyno := c.PostForm("applyno")
	var temp=""
	session := sessions.Default(c)
	var error error
	if(session.Get("flag").(string)=="m"){
		error=db.MyDB.QueryRow("exec UndoManagerCheck @applyno=?",applyno).Scan(&temp)
	}else {
		error=db.MyDB.QueryRow("exec UndoCheck @applyno=?",applyno).Scan(&temp)
	}
	if(temp!=tno){
		fmt.Println("tno与applyno的tno不相等",temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		_,error:=db.MyDB.Exec("exec Undo @applyno=?",applyno)
		if error == nil {
			c.JSON(http.StatusOK, "ok")
			fmt.Println("删除成功")
		}else {
			c.JSON(http.StatusOK, error)
			fmt.Println("删除失败",error)
		}
	}
}
func MangerUndo(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	applyno := c.PostForm("applyno")
	var temp=""
	error:=db.MyDB.QueryRow("exec UndoCheck @applyno=?",applyno).Scan(&temp)
	if(temp!=mno){
		fmt.Println("tno与applyno的tno不相等",temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		_,error:=db.MyDB.Exec("exec Undo @applyno=?",applyno)
		if error == nil {
			c.JSON(http.StatusOK, "ok")
			fmt.Println("删除成功")
		}else {
			c.JSON(http.StatusOK, error)
			fmt.Println("删除失败",error)
		}
	}
}
