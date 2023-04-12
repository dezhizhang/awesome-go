package router

import (
	"bluebell/controller"
	"bluebell/middlerware"
	"github.com/gin-gonic/gin"
)

func StepRouter() *gin.Engine {
	r := gin.New()

	g := r.Group("/api/v1")

	g.POST("/register", controller.RegisterHandler)

	g.POST("/login", controller.LoginHandler)

	g.GET("/user", middlerware.Auth(), controller.UserByIdHandler)
	g.GET("/user/list", middlerware.Auth(), controller.UserListHandler)

	return r
}
