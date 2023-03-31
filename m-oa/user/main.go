package main

import (
	srv "common"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	srv.Run(r, "user", ":8081")
}
