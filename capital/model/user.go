package model

//用户表

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"'`
	Email    string `json:"email"`
}
