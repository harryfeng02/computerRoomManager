package apis

import (
	. "computerRoomManager/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"path"
)

func LoginApi(c *gin.Context) {
	var yzm_map map[string]string
	yzm_map = make(map[string]string)
	yzm_map["1.png"] = "7364"
	yzm_map["2.jpg"] = "46168"

	userno := c.PostForm("usrno")
	passwd := c.PostForm("passwd")
    yzm_value:=c.PostForm("yzm")
	yzm_name:=c.PostForm("yzmname")

	if(yzm_map[path.Base(yzm_name)]!=yzm_value){
		c.JSON(http.StatusOK, gin.H{
			"code":  202,
		})
		return
	}

	manager := Manager{Mno: userno, Mname: "", Mphone: "", Mpassword: passwd}
	name, err := manager.Check()
	if err != nil {
		//log.Fatal(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":  201,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"name": name,
	})
	//c.Redirect(http.StatusMovedPermanently, "index.html")
}
