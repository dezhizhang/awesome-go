package controller

import (
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/contact.html"))
	must.Execute(w, "")
}
