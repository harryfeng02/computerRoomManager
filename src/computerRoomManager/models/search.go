package models

import (
	db "computerRoomManager/database"
	"fmt"
)

type SearchRoom struct {
	Location string
	SfName    string
	StartTime    string
	EndTime    string
	StratLesson string
	EndLesson string
	PeopleNum int
}
type resultJson struct {
	Cpno string
	Date string
	Location    string
	Building  string
	CpName      string
	Peoplenum int
	Flag string
}

func (require *SearchRoom) Search() (rooms []resultJson, err error) {
	rows, err := db.MyDB.Query(
		"select computerroomstatus.cpno,cplocation,cpbuilding,cpname,computerroomstatus.date,flag,computerroomstatus.peoplenum from computerroom,computerroomstatus where  computerroomstatus.cpno in (select cpno from own where sfname like ?)  and computerroomstatus.cpno=computerroom.cpno and ?<=computerroomstatus.date and computerroomstatus.date<=? and computerroomstatus.peoplenum>=? and flag>=? and (flag+1)<=? and computerroom.cplocation=?",
		"%"+require.SfName+"%",require.StartTime,require.EndTime, require.PeopleNum,require.StratLesson,require.EndLesson,require.Location)
	if err != nil {
		return nil, err
	} else {
		var resultRooms []resultJson
		for rows.Next() {
			var resultRoom resultJson
			if err := rows.Scan(&resultRoom.Cpno,&resultRoom.Location, &resultRoom.Building, &resultRoom.CpName, &resultRoom.Date, &resultRoom.Flag, &resultRoom.Peoplenum); err == nil {
				resultRooms = append(resultRooms, resultRoom)
				fmt.Println(resultRoom.Building)
			}
		}
		return resultRooms, err
	}
}
