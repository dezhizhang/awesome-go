package router

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
)

func StepRouter() *gin.Engine {
	r := gin.New()

	g := r.Group("/api/v1")

	g.POST("/register", controller.RegisterHandler)

	g.GET("/user", func(c *gin.Context) {

	})

	return r
}
