package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "发送的地址：", r.URL.Path, r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}
}
