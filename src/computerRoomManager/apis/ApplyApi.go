package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "computerRoomManager/database"
	"fmt"
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

	err1:=db.MyDB.QueryRow("select mname,mphone from manager,manage where manager.mno=manage.mno and manage.cpno=? and manage.date=?", cpno,date).
		Scan(&mname,&mphone)
	result,err2:=db.MyDB.Query("select sfname,sfverison from own where cpno=?", cpno)
	class,err3:=db.MyDB.Query("  select tid,cname,teach.cclass,periods_remain from teach,course where tno=? and teach.cno=course.cno", tno)
	fmt.Println("err1",err1)
	fmt.Println("err2",err2)
	fmt.Println("err3",err3)
	if err2 != nil || err1!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err2,
		})
		return
	} else {
		var temp string
		for result.Next() {
			var sf SoftWare
			if err := result.Scan(&sf.SfName,&sf.SV); err == nil {
				temp=temp+sf.SfName+sf.SV+";"
			}
		}
		var temp2[] Class
		for class.Next() {
			var c Class
			if err := class.Scan(&c.Tid,&c.Cname,&c.Cclass,&c.Periods_remain); err == nil {
				temp2=append(temp2,c)
				fmt.Println(c.Periods_remain)
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
//func ApplyRoom(c *gin.Context) {
//	tno:=isHaveSessions(c)
//	if(tno==""){
//		return
//	}
//	tid:=c.PostForm("tid")
//	date:=c.PostForm("date")
//	flag:=c.PostForm("flag")
//	cpno:c.PostForm("cpno")
//	var mname string
//	var mphone string
//
//	err1:=db.MyDB.QueryRow("select mname,mphone from manager,manage where manager.mno=manage.mno and manage.cpno=? and manage.date=?", cpno,date).
//		Scan(&mname,&mphone)
//	result,err2:=db.MyDB.Query("select sfname,sfverison from own where cpno=?", cpno)
//	class,err3:=db.MyDB.Query("  select tid,cname,teach.cclass,periods_remain from teach,course where tno=? and teach.cno=course.cno", tno)
//	fmt.Println("err1",err1)
//	fmt.Println("err2",err2)
//	fmt.Println("err3",err3)
//	if err2 != nil || err1!=nil {
//		c.JSON(http.StatusOK,gin.H{
//			"code":"201",
//			"error":err2,
//		})
//		return
//	} else {
//		var temp string
//		for result.Next() {
//			var sf SoftWare
//			if err := result.Scan(&sf.SfName,&sf.SV); err == nil {
//				temp=temp+sf.SfName+sf.SV+";"
//			}
//		}
//		var temp2[] Class
//		for class.Next() {
//			var c Class
//			if err := class.Scan(&c.Tid,&c.Cname,&c.Cclass,&c.Periods_remain); err == nil {
//				temp2=append(temp2,c)
//				fmt.Println(c.Periods_remain)
//			}
//		}
//		c.JSON(http.StatusOK,gin.H{
//			"code":"200",
//			"mname":mname,
//			"mphone":mphone,
//			"sfware":temp,
//			"cclass":temp2,
//		})
//	}
//}

