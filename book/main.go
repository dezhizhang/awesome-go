package main

import (
	"html/template"
	"log"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//	r.Header.Set("Content-Type", "application/json")
//	user := model.User{
//		Id:       "123",
//		Name:     "小明",
//		Password: "123456",
//		Email:    "1541609448@qq.com",
//	}
//	marshal, err := json.Marshal(&user)
//	if err != nil {
//		log.Printf("序列化失败%s", err)
//	}
//	w.Write(marshal)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Printf("加载模板失败%s", err)
	}
	files.Execute(w, nil)
}

func main() {
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Printf("启动服务失败%s", err)
	}

}
