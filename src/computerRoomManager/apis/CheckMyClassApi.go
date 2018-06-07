package apis
import (
	db "computerRoomManager/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func CheckMyClass(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	var cname string
	var cno string
	var cclass string
	var location  string
	var building  string
	var peoplenum  string
	var time  string
	var semester  string
	var credit  string
	pp:=[2]gin.H{}

	dir:=db.MyDB.QueryRow("select cname,course.cno,course.cclass,location,building,semester,time,peoplenum,credit from teach,course where course.cno=teach.cno and teach.tno=?", tno).
		Scan(&cname,&cno,&cclass,&location,&building,&semester,&time,&peoplenum,&credit)
	fmt.Println(cclass,dir)
	pp[0]= gin.H{
		"course":cname,
		"course_code": cno+"-"+cclass,
		"time":time,
		"location": location,
		"building":building,
		"semester":semester,
		"peoplenum": peoplenum,
		"credit":credit,
	}
	pp[1]= gin.H{
		"course":cname,
		"course_code": cno+"-"+cclass,
		"time":time,
		"location": location,
		"building":building,
		"semester":semester,
		"peoplenum": peoplenum,
		"credit":credit,
	}
	c.JSON(http.StatusOK, pp)
}