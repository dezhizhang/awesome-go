package main

import (
	"book/controller"
	"log"
	"net/http"
)

func main() {

	// 加载静态资源
	server := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", server))

	// 路由请求
	http.HandleFunc("/", controller.HomeHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/contact", controller.ContactHandler)
	http.HandleFunc("/cart", controller.CartHandler)

	//-------------
	//测试文件
	http.HandleFunc("/test", controller.TestHandler)

	baseAPI := "/api/v1"

	http.HandleFunc(baseAPI+"/user/check-email", controller.CheckEmail)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
