package models

import (
	db "computerRoomManager/database"
)

type SearchRoom struct {
	SfName    string
	PeopleNum int
}

func (require *SearchRoom) Search() (rooms []Computeroom, err error) {
	rows, err := db.MyDB.Query(
		"select * from dbo.computerroom where cpno IN (select cpno from dbo.own where sfname=? )and peoplenum >= ? ",
		require.SfName, require.PeopleNum)
	if err != nil {
		return nil, err
	} else {
		var resultRooms []Computeroom
		for rows.Next() {
			var resultRoom Computeroom
			if err := rows.Scan(&resultRoom.CpNo, &resultRoom.CpName, &resultRoom.PeopleNum); err == nil {
				resultRooms = append(resultRooms, resultRoom)
			}
		}
		return resultRooms, err
	}
}
