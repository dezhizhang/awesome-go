package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/index.html"))
	must.Execute(w, "")
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/login.html"))
	must.Execute(w, "")
}

func main() {

	// 加载静态资源
	server := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", server))

	// 路由请求
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", handlerLogin)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
