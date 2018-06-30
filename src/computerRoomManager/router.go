package main

import (
	. "computerRoomManager/apis"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	//store.Options(sessions.Options{MaxAge: 3600})
	router.Use(sessions.Sessions("mysession", store))

	//dir2:="C://Users//bingnan//computerRoomManager//src//computerRoomManager"
	dir2:="C://Users//bingnan//computerRoomManager//src//computerRoomManager"
	router.LoadHTMLGlob(filepath.Join(dir2,"htmls/*"))
	//登录界面
	router.GET("/", PreLoginApi)
	router.POST("/login", LoginApi)
	router.POST("/search", SearchApi)

	router.GET("/index.html", IndexApi)
	router.POST("/room", AddComRoomApi)
	router.GET("/home.html", HomeInitApi)
	router.GET("/manage.html", MangerHomeInitApi)
	router.GET("/checkApply.html", MangerCheckApi)
	router.GET("/checkSoftware.html", MangerUpdateApi)
	router.GET("/detail.html", DetailApi)
	router.GET("/yzmchange", YzmChnageapi)
	router.GET("/check.html", CheckApi)

	router.GET("/check",CheckApply)
	router.GET("/myclass.html",MyClassApi)
	router.GET("/myclass",CheckMyClass)
	router.POST("/undo",Undo)
	router.POST("/checkinfo",CheckInfo)
	router.POST("/applyroom",ApplyRoom)
	router.POST("/lookmore",LookMore)
	router.GET("/logout",LogoutApi)
	router.GET("/checkm",ManagerCheckApply)
	router.POST("/lookmorem",ManagerLookMore)
	router.GET("/checkSoftware",CheckSoftWare)
	router.POST("/deletesoftware",DeleteSoftware)
	router.POST("/updatesoftware",UpdateSoftware)
	router.POST("/newsoftware",NewSoftware)
	router.GET("/getpeoplenum",Checkpeoplenum)
	router.POST("/updatepeoplenum",Updatepeoplenum)
	return router
}
