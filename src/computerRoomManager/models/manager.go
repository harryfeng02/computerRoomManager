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

func (manage *Manager) Check() (name string, err error) {
	no := manage.Mno
	passwd := manage.Mpassword
	var username string
	err2 := db.MyDB.QueryRow("select dbo.manager.mname from dbo.manager where dbo.manager.mno=? and dbo.manager.mpassword=?",
		no, passwd).Scan(&username)
	if err2 == nil {
		return username, err2
	} else {
		return "", err2
	}
}
