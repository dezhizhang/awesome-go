package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自定义请求")
}
func main() {
	handler := MyHandler{}
	http.Handle("/hello", &handler)
	http.ListenAndServe("localhost:8080", nil)
}
