package controller

import (
	"html/template"
	"net/http"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("./view/cart.html"))
	must.Execute(w, "")
}
