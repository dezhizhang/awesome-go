package service

import (
	"book/dao"
	"book/model"
)

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
