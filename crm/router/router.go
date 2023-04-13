package router

import (
	"crm/controller"
	"crm/middlerware"
	"github.com/gin-gonic/gin"
)

func StepRouter() *gin.Engine {
	r := gin.New()

	// 注册与登录
	r.POST("/api/v1/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	g := r.Group("/api/v1")
	g.Use(middlerware.Auth())

	g.GET("/user", controller.UserByIdHandler)
	g.GET("/user/list", controller.UserListHandler)

	//客户管理
	g.GET("/customer/contact", controller.ContactHandler)

	return r
}
