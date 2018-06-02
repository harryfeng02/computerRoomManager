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

type resultJson struct {
	Id        int
	Xiaoqu    string
	Building  string
	Cpno      string
	Peoplenum int
	IdTwo     int
}

func SearchApi(c *gin.Context) {
	//pageSize := c.PostForm("pageSize")
	//pageNumber := c.PostForm("pageNumber")
	xiaoqu := c.PostForm("xiaoqu")
	software := c.PostForm("software")
	//time_start := c.PostForm("time_start")
	//time_end := c.PostForm("time_end")
	peoplenum := c.PostForm("peoplenum")

	number, _ := strconv.Atoi(peoplenum)
	require := SearchRoom{SfName: software, PeopleNum: number}

	results, err := require.Search()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	var resultjsons []resultJson
	for i := 0; i < len(results); i++ {
		resultjson := resultJson{
			Id:        i + 1,
			Xiaoqu:    xiaoqu,
			Cpno:      results[i].CpNo,
			Building:  results[i].CpName,
			Peoplenum: results[i].PeopleNum,
			IdTwo:     i + 1,
		}
		resultjsons = append(resultjsons, resultjson)
	}

	c.JSON(http.StatusOK, gin.H{
		"total": len(results),
		"rows":  resultjsons[0],
	})

}
