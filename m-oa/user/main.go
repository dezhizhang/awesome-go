package main

import (
	"awesome-go/m-oa/user/config"
	"awesome-go/m-oa/user/router"
	srv "common"
	"common/logs"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()

	lc := &logs.LogConfig{
		DebugFileName: "logs/debug.log",
		InfoFileName:  "logs/info.log",
		WarnFileName:  "logs/warn.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}

	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

	router.InitRouter(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
