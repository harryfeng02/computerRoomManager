package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "computerRoomManager/database"
	"fmt"
	"strconv"
)
type SoftWare struct {
	SfName string
	SV string
}
type Class struct {
	Tid string
	Cname string
	Cclass string
	Periods_remain string
}

func CheckInfo(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	cpno:=c.PostForm("cpno")
	date:=c.PostForm("date")

	var mname string
	var mphone string

	err1:=db.MyDB.QueryRow("exec LookManager @cpno=?,@date=?", cpno,date).
		Scan(&mname,&mphone)
	result,err2:=db.MyDB.Query("exec LookSoftWare @cpno=?", cpno)
	class,err3:=db.MyDB.Query("exec LookMyClass @tno=?", tno)

	if  err1!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err1,
		})
		fmt.Println("管理员搜索error",err1)
		return
	} else if err2 != nil  {
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err2,
		})
		fmt.Println("机房环境搜索error",err2)
		return
	} else if err3 != nil  {
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err3,
		})
		fmt.Println("授课搜索error",err3)
		return
	}else{
		var temp string
		for result.Next() {
			var sf SoftWare
			if err := result.Scan(&sf.SfName,&sf.SV); err == nil {
				temp=temp+sf.SfName+" "+sf.SV+";"
			}
		}
		var temp2[] Class
		for class.Next() {
			var c Class
			if err := class.Scan(&c.Tid,&c.Cname,&c.Cclass,&c.Periods_remain); err == nil {
				temp2=append(temp2,c)
			}
		}
		c.JSON(http.StatusOK,gin.H{
			"code":"200",
			"mname":mname,
			"mphone":mphone,
			"sfware":temp,
			"cclass":temp2,
		})
	}
}
func ApplyRoom(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	tid:=c.PostForm("tid")
	date:=c.PostForm("date")
	flag:=c.PostForm("flag")
	cpno:=c.PostForm("cpno")

	flag2,_ := strconv.Atoi(flag)
	tid2,_ := strconv.Atoi(tid)
	flag2=flag2+1

	var remain string
    var peoplenum int

	err1:=db.MyDB.QueryRow("exec LookPeriods_Remain @tno=?,@tid=?", tno,tid2).
		Scan(&remain,&peoplenum)
	var peoplenum2 int
	err3:=db.MyDB.QueryRow("exec LookPeople @cpno=?,@falg=?,@date=?", cpno,flag,date).
		Scan(&peoplenum2)
	if(err1!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err1,
		})
		fmt.Println("剩余课时搜索失败",err1)
		return
	}else if(err3!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"204",
			"error":err3,
		})
		fmt.Println("机房剩余人数搜索失败",err3)
		return
	}else if(peoplenum2<peoplenum){
		c.JSON(http.StatusOK,gin.H{
			"code":"204",
			"error":"人数已满",
		})
		return
	}else if(remain=="0"){
		c.JSON(http.StatusOK,gin.H{
			"code":"202",
		    "error":"剩余课时不足",
		})
		return
	}else{
       SQL:="insert into apply values(?,'"+date+"','',(select CONVERT(varchar,GETDATE(),20)),?,?,?,?)"
		_, err2 :=db.MyDB.Exec(SQL,cpno,flag2-1,flag2,peoplenum,tid2)
		if(err2!=nil){
         c.JSON(http.StatusOK,gin.H{
			"code":"203",
			"error":err2,
		})
		fmt.Println("插入失败",err2)
         return
        }
        c.JSON(http.StatusOK,gin.H{
			"code":"200",
			"error":"申请成功",
		})
		return
	}
}

