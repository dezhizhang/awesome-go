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

// 注册用户

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

// 用户登录

func UserLogin(req *model.UserLogin) (err error) {
	var user model.User
	_, err = CheckUserExit(req.Email)
	if err != nil {
		return err
	}
	result := dao.DB.Where("email = ? and password = ? ", req.Email, req.Password).Find(&user)
	return result.Error
}
