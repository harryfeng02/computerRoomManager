package apis

import (
	. "computerRoomManager/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComRoomApi(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	roomCpNo := c.Request.FormValue("CpNo")
	roomName := c.Request.FormValue("CpName")
	roomPeoNum := c.Request.FormValue("PeopleNum")

	num, err := strconv.Atoi(roomPeoNum)
	if err != nil {
		log.Fatal(err.Error())
	}
	room := Computeroom{CpNo: roomCpNo, CpName: roomName, PeopleNum: num}

	err = room.AddComputeRoom()
	if err != nil {
		log.Fatal(err.Error())
	}

	//msg := fmt.Sprintf("insert successful")
	c.Redirect(http.StatusMovedPermanently, "index.html")
}
