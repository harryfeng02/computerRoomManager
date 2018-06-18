package models

import (
	db "computerRoomManager/database"
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
		"exec SearchRoom @sfname=?,@startdate=?,@enddate=?,@peoplenum=?,@startlesson=?,@endlesson=?,@location=?",
		require.SfName,require.StartTime,require.EndTime, require.PeopleNum,require.StratLesson,require.EndLesson,require.Location)
	if err != nil {
		return nil, err
	} else {
		var resultRooms []resultJson
		for rows.Next() {
			var resultRoom resultJson
			if err := rows.Scan(&resultRoom.Cpno,&resultRoom.Location, &resultRoom.Building, &resultRoom.CpName, &resultRoom.Date, &resultRoom.Flag, &resultRoom.Peoplenum); err == nil {
				resultRooms = append(resultRooms, resultRoom)
			}
		}

		return resultRooms, err
	}
}
