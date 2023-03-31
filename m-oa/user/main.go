package main

import (
	"awesome-go/m-oa/user/router"
	srv "common"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.InitRouter(r)

	srv.Run(r, "user", ":8081")
}
