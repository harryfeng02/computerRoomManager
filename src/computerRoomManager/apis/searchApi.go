package apis

import (
	"strconv"
	. "computerRoomManager/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	datetime = "2006-01-02 15:04:02"
)



func SearchApi(c *gin.Context) {
	tno:=isHaveSessions(c)
	if(tno==""){
		return
	}
	//pageSize := c.PostForm("pageSize")
	//pageNumber := c.PostForm("pageNumber")
	location := c.PostForm("location")
	software := c.PostForm("software")
	time_start := c.PostForm("time_start")
	time_end := c.PostForm("time_end")
	lesson_start:=c.PostForm("lesson_start")
	lesson_end:=c.PostForm("lesson_end")
	peoplenum := c.PostForm("peoplenum")

	number, _ := strconv.Atoi(peoplenum)
	require := SearchRoom{SfName: software, PeopleNum: number,StartTime:time_start,EndTime:time_end,StratLesson:lesson_start,EndLesson:lesson_end,Location:location}
	results, err := require.Search()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "201",
			"error": err.Error(),
		})
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"rows":  results,
		"total":len(results),
	})

}

