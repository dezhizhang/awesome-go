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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/login.html"))
	must.Execute(w, "")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/register.html"))
	must.Execute(w, "")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/contact.html"))
	must.Execute(w, "")
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/cart.html"))
	must.Execute(w, "")
}

func main() {

	// 加载静态资源
	server := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", server))

	// 路由请求
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/cart", cartHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
