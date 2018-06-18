package models

import (
	db "computerRoomManager/database"
)


type Manager struct {
	Mno       string
	Mname     string
	Mphone    string
	Mpassword string
}

func (manage *Manager) MCheck() (name string, err error) {
	no := manage.Mno
	passwd := manage.Mpassword
	var username string
	err2 := db.MyDB.QueryRow("exec isManagerPasswordRight @mno=?,@password=?",
		no, passwd).Scan(&username)
	if err2 == nil {
		return username, err2
	} else {
		return "", err2
	}
}
