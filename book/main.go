package main

import (
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("./view/index.html")
	if err != nil {
		log.Printf("解析模板失败%s", err)
	}
	file.Execute(w, "周华建")
}

func main() {

	// 加载静态资源
	server := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", server))

	// 路由请求
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
