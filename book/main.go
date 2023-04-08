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
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
