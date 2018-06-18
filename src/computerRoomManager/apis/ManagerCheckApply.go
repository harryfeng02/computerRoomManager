package apis
import (
	db "computerRoomManager/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func ManagerCheckApply(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}

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
	var apply_result_json =[]gin.H{}

	apply_result,err:=db.MyDB.Query("exec LookApplication @mno=?",mno)

	if(err!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err,
		})
		fmt.Println("管理员查看申请error：",err)
		return
	}

	for apply_result.Next() {
		if err := apply_result.Scan(&applyno,&cname,&cno,&cclass,&starttime,&start_lesson,&end_lesson,&cplocation,&cpbuilding,&cpname,&peoplenum); err == nil {
			temp:=gin.H{
				"apply":  applyno,
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

func ManagerLookMore(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	applyno:=c.PostForm("applyno")

	var tname string
	var tphone string
	var applytime string

	err1:=db.MyDB.QueryRow(" exec ManagerLookMore @applyno=?", applyno).
		Scan(&applytime,&tname,&tphone)

	if(err1!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err1,
		})
		fmt.Println("查看教师申请error：",err1)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":"200",
		"tname":tname,
		"tphone":tphone,
		"applytime":applytime,
	})
}