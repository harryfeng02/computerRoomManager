package apis
import (
	db "computerRoomManager/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func CheckApply(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}

	var applyno string
	var cpno  int
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
	var apply_result_json =[]gin.H{}

	fmt.Println(time.Now())
	apply_result,err:=db.MyDB.Query("exec checkApply @tno=?",tno)
	fmt.Println(time.Now())

	if(err!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err,
		})
		fmt.Println("查看申请error：",err)
		return
	}

	for apply_result.Next() {
		if err := apply_result.Scan(&applyno,&cpno,&cname,&cno,&cclass,&starttime,&start_lesson,&end_lesson,&cplocation,&cpbuilding,&cpname,&peoplenum); err == nil {
			temp:=gin.H{
				"apply":  applyno,
				"cpno":cpno,
				"course":cname,
				"course_code": cno+"-"+cclass,
				"date": starttime,
				"time":start_lesson+"-"+end_lesson+"节",
				"xiaoqu": cplocation,
				"location": cpbuilding+cpname,
				"peoplenum": peoplenum,
			}
			apply_result_json =append(apply_result_json,temp)
		}else {
			fmt.Println("查看申请scan error：",err)
		}
	}
	c.JSON(http.StatusOK, apply_result_json)
}
func LookMore(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	cpno:=c.PostForm("cpno")
	date:=c.PostForm("date")

	var mname string
	var mphone string

	err1:=db.MyDB.QueryRow("exec Lookmore @cpno=?,@date=?", cpno,date).
		Scan(&mname,&mphone)
	sf_result,err2:=db.MyDB.Query("exec LookSoftWare @cpno=?", cpno)

	if(err1!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err1,
		})
		fmt.Println("查看管理员error：",err1)
		return
	}else if(err2!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err2,
		})
		fmt.Println("查看环境error：",err2)
		return
	}
	var temp string
	for sf_result.Next() {
		var sf SoftWare
		if err := sf_result.Scan(&sf.SfName,&sf.SV); err == nil {
			temp=temp+sf.SfName+ " "+sf.SV+";"
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"mname":mname,
		"mphone":mphone,
		"sfware":temp,
	})
}