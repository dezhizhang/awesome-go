package controller

import (
	"book/service"
	"html/template"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/login.html"))
	must.Execute(w, "")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/register.html"))
	must.Execute(w, "")
}

func Login(w http.ResponseWriter, r *http.Request) {
	//email := r.PostFormValue("email")
	//password := r.PostFormValue("password")
	//_, err := service.CheckUserIsExist(email, password)
	//if err != nil {
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}
	//http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	_, err := service.CheckUserIsExist(email)
	if err != nil {
		log.Printf("邮箱已存在%s", err)
		w.Write([]byte("邮箱已存在"))
		return
	}
	w.Write([]byte("邮箱验证成功"))
}
