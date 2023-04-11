package logic

import (
	"bluebell/dao"
	"bluebell/model"
	"log"
)

func CheckUserExit(email string) (user *model.User, err error) {
	err = dao.DB.Find(&user).Where("email =? ", email).Error
	if err != nil {
		log.Printf("查询失败%s", err)
		return nil, err
	}
	return user, nil
}
