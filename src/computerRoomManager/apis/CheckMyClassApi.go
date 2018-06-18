package apis
import (
	db "computerRoomManager/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)
type MyClass struct {
	Course    string
	Course_No string
	Course_Class string
	Time      string
	Location  string
	Building  string
	Semester  string
	Peoplenum string
	Credit    string
}
func CheckMyClass(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	fmt.Println(tno)
	myclass_result,err1:=db.MyDB.Query("exec CheckMyClass @tno=?", tno)
	
	if(err1!=nil){
        c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err1,
		})
        fmt.Println("查看授课error：",err1)
		return

	} else{
		var temp []MyClass
		for myclass_result.Next() {
			var mc MyClass
			if err := myclass_result.Scan(&mc.Course,&mc.Course_No,&mc.Course_Class,&mc.Location,&mc.Building,&mc.Semester,&mc.Time,&mc.Peoplenum,&mc.Credit); err == nil {
				temp=append(temp,mc)
			}else {
				fmt.Println("查看授课scan error：",err)
			}
		}
		c.JSON(http.StatusOK, temp)
	}
}