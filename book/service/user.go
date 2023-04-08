package service

import (
	"book/dao"
	"book/model"
	"log"
)

// 检测用户是否存在

func CheckUserIsExist(email string, password string) (*model.User, error) {
	var user model.User
	sqlStr := "select * from user where email = ? and password = ? "
	row := dao.DB.QueryRow(sqlStr, email, password)
	err := row.Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 添加用户

func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	_, err := dao.DB.Exec(sqlStr, username, password, email)
	if err != nil {
		log.Printf("添加用户失败%s", err)
		return err
	}
	return nil
}
