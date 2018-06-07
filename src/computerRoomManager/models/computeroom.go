package models

import (
	db "computerRoomManager/database"
)

type Computeroom struct {
	CpNo   string `json:"CpNo"`
	CpName string `json:"CpName"`
	//如果有中文  golang默认编码为utf-8 sql server中应设置nvarchar而不是varchar才会转换为utf-8
	PeopleNum int `json:"PeopleNum"`
	CpLocation string
	CpBuilding string
	date string
	flag int
}

//添加机房
func (room *Computeroom) AddComputeRoom() (err error) {
	tx, _ := db.MyDB.Begin()
	defer tx.Commit()
	//room.CpNo, room.CpName, room.PeopleNum
	_, err2 := tx.Exec("insert into dbo.computerroom values(?,?,?)", room.CpNo, room.CpName, room.PeopleNum)
	if err2 != nil {
		return err2
	}
	return nil
}

func Query(command string) (rooms []Computeroom, err error) {
	pre2, err := db.MyDB.Prepare(command)
	if err != nil {
		return nil, err
	}
	rows, err := pre2.Query()
	if err != nil {
		return nil, err
	} else {
		var rooms []Computeroom
		for rows.Next() {
			var room Computeroom
			if err := rows.Scan(&room.CpNo, &room.CpName, &room.PeopleNum); err == nil {
				rooms = append(rooms, room)
			}
		}
		return rooms, nil
	}
}
