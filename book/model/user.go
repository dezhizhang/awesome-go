package model

//用户表

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"'`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
