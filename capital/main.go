package main

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/index.html"))
	must.Execute(w, "")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/login.html"))
	must.Execute(w, "")
}

func main() {

	// 处理静态资源
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", handleLogin)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
