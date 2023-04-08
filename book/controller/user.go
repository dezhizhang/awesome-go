package controller

import (
	"book/service"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	_, err := service.CheckUserIsExist(email, password)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
