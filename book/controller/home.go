package controller

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/index.html"))
	must.Execute(w, "")
}
