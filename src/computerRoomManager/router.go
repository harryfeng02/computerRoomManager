package main

import (
	. "computerRoomManager/src/computerRoomManager/apis"

	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"path/filepath"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	dir2:="D://GoPkg//src//computerRoomManager//src//computerRoomManager"
	router.LoadHTMLGlob(filepath.Join(dir2,"htmls/*"))
	fmt.Println(GetCurrentDirectory())
	//登录界面
	router.GET("/", PreLoginApi)
	router.POST("/login", LoginApi)
	router.POST("/search", SearchApi)

	router.GET("/index.html", IndexApi)
	router.POST("/room", AddComRoomApi)
	router.GET("/home.html", HomeInitApi)

	router.GET("/detail.html", DetailApi)
	router.GET("/yzmchange", YzmChnageapi)
	router.GET("/check.html", CheckApi)

	router.GET("/test.html",WhatApi)
	return router
}
func GetCurrentDirectory() string {
	dir,_:=os.Getwd()
	return dir
}