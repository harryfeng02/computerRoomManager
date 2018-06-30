package apis

import ("github.com/gin-gonic/gin"
db "computerRoomManager/database"
	"net/http"
	"fmt"
)

func CheckSoftWare(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	var cpno string
	var sfname string
	var sfverison string
	sf_result,err:=db.MyDB.Query("exec LookSoftwareEnvironment @mno=?",mno)
	if(err!=nil){
		c.JSON(http.StatusOK,gin.H{
			"code":"201",
			"error":err,
		})
		fmt.Println("管理员查看软件环境error：",err)
		return
	}
	var sf_result_json =[]gin.H{}
	for sf_result.Next() {
		if err := sf_result.Scan(&cpno,&sfname,&sfverison); err == nil {
			temp:=gin.H{
				"cpno":cpno,
				"sfname":  sfname,
				"sfverison":sfverison,
			}
			sf_result_json =append(sf_result_json,temp)
		}else {
			fmt.Println("查看申请scan error：",err)
		}
	}
	c.JSON(http.StatusOK, sf_result_json)
}

func DeleteSoftware(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	cpno := c.PostForm("cpno")
	sfname := c.PostForm("sfname")
	sfverison := c.PostForm("sfversion")
	var temp=""
	error:=db.MyDB.QueryRow("exec DeleteSFCheck @cpno=?",cpno).Scan(&temp)
	if(temp!=mno||error!=nil){
		fmt.Println("mno与manage的mno不相等",temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		_,error:=db.MyDB.Exec("exec DeleteSoftwareEnvironment @cpno=?,@sfname=?,@sfversion=?",cpno,sfname,sfverison)
		if error == nil {
			c.JSON(http.StatusOK, "ok")
			fmt.Println("删除成功")
		}else {
			c.JSON(http.StatusOK, error)
			fmt.Println("删除失败",error)
		}
	}
}
func UpdateSoftware(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	cpno := c.PostForm("cpno")
	sfname := c.PostForm("sfname")
	sfverison := c.PostForm("sfversion")
	oldsfname:=c.PostForm("oldsfname")
	oldsfversion:=c.PostForm("oldsfversion")
	var temp=""
	error:=db.MyDB.QueryRow("exec DeleteSFCheck @cpno=?",cpno).Scan(&temp)
	if(temp!=mno||error!=nil){
		fmt.Println("mno与manage的mno不相等",temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		_,error:=db.MyDB.Exec("exec UpdateSoftware @cpno=?,@oldsfname=?,@oldsfversion=?,@sfname=?,@sfversion=?",cpno,oldsfname,oldsfversion,sfname,sfverison)
		if error == nil {
			c.JSON(http.StatusOK, "ok")
			fmt.Println("更新成功")
		}else {
			c.JSON(http.StatusOK, error)
			fmt.Println("更新失败",error)
		}
	}
}
func NewSoftware(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	cpno := c.PostForm("cpno")
	sfname := c.PostForm("sfname")
	sfverison := c.PostForm("sfversion")
	var temp=""
	error:=db.MyDB.QueryRow("exec DeleteSFCheck @cpno=?",cpno).Scan(&temp)
	if(temp!=mno||error!=nil){
		fmt.Println("mno与manage的mno不相等",temp,error)
		c.JSON(http.StatusOK, "false")
	}else {
		_,error:=db.MyDB.Exec("exec AddSoftwareEnvironment @cpno=?,@sfname=?,@sfversion=?",cpno,sfname,sfverison)
		if error == nil {
			c.JSON(http.StatusOK, "ok")
			fmt.Println("添加成功")
		}else {
			c.JSON(http.StatusOK, error)
			fmt.Println("添加失败",error)
		}
	}
}

func Checkpeoplenum(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	var num int
	error:=db.MyDB.QueryRow("exec Checkpeoplenum @mno=?",mno).Scan(&num)
	if(error!=nil){
		fmt.Println("查询机房人数失败",error)
		c.JSON(http.StatusOK, gin.H{
			"code":"error",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":"ok",
			"num":  num,
		})
		fmt.Println("查询机房人数成功",error)
	}
}

func Updatepeoplenum(c *gin.Context) {
	mno:=isHaveSessions(c)
	if(mno==""){
		return
	}
	num:=c.PostForm("newpeoplenum")
	_,error:=db.MyDB.Exec("exec Updatepeoplenum @mno=?,@newpeoplenum=?",mno,num)
	if(error!=nil){
		fmt.Println("更新机房人数失败",error)
		c.JSON(http.StatusOK, gin.H{
			"code":"error",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":"ok",
		})
		fmt.Println("更新机房人数成功",error)
	}
}
