package main

import (
	"time"
	"fmt"
)

func test()  {
	hh,_:=time.ParseDuration("24h")
	//var sz [310]string
	now:=time.Now()
	for i := 0; i < 31; i++ {
		now=now.Add(hh)
		for j := 0; j < 8; j++ {
			fmt.Println(now.Format("2006-01-02")+"")
		}
	}
}