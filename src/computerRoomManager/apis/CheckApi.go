package apis
import (
	db "computerRoomManager/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func CheckApply(c *gin.Context) {
	var applyno string
	var cname string
	var cno string
	var cclass string
	var starttime  string
	var start_lesson  string
	var end_lesson  string
	var cplocation  string
	var cpbuilding  string
	var cpname  string
	var peoplenum  string
	pp:=[2]gin.H{}

	dir:=db.MyDB.QueryRow("SELECT applyno,cname,course.cno,course.cclass,starttime,start_lesson,end_lesson,cplocation,cpbuilding,cpname,apply.peoplenum FROM Apply,teach, computerroom,course where apply.tid in (select tid from teach where teach.tno=?)  and apply.cpno=computerroom.cpno and teach.tid in (select tid from teach where teach.tno=?) and course.cno in (select cno from teach where teach.tno=?)", "000000", "000000","000000").
		Scan(&applyno,&cname,&cno,&cclass,&starttime,&start_lesson,&end_lesson,&cplocation,&cpbuilding,&cpname,&peoplenum)
	fmt.Println(cclass,dir)
	pp[0]= gin.H{
		"apply":  applyno,
		"course":cname,
		"course_code": cno+"-"+cclass,
		"date": starttime,
		"time":start_lesson+"-"+end_lesson+"节",
		"xiaoqu": cplocation,
		"location": cpbuilding+cpname,
		"peoplenum": peoplenum,
	}
	pp[1]= gin.H{
		"apply":  applyno,
		"course":cname,
		"course_code": cno+"-"+cclass,
		"date": starttime,
		"time":start_lesson+"-"+end_lesson+"节",
		"xiaoqu": cplocation,
		"location": cpbuilding+cpname,
		"peoplenum": peoplenum,
	}
	c.JSON(http.StatusOK, pp)
}