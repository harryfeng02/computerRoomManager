package models

import (
	db "computerRoomManager/database"
)
func  Check(tno string,tpwd string) (name string, err error) {
	var userid string
	err2 := db.MyDB.QueryRow("exec isTeacherPasswordRight @tno=?,@tpassword=?",
		tno, tpwd).Scan(&userid)
	if err2 == nil {
		return userid, err2
	} else {
		return "", err2
	}
}
