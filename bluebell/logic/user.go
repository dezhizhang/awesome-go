package logic

import (
	"bluebell/dao"
	"bluebell/model"
	"bluebell/utils"
	"log"
	"time"
)

func CheckUserExit(email string) (user *model.User, err error) {

	err = dao.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Printf("查询失败%s", err)
		return nil, err
	}
	return user, nil
}

func CreateUser(req *model.UserRegister) (err error) {
	var user model.User
	id, err := utils.SnowflakeId()
	if err != nil {
		log.Printf(err.Error())
	}
	now := time.Now()
	user.Id = id
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password
	user.CreateTime = now.Format("2006-01-02 15:04:05")
	err = dao.DB.Create(user).Error
	return
}
