package logic

import (
	"bluebell/dao"
	"bluebell/model"
	"bluebell/utils"
	"fmt"
	"log"
	"time"
)

const FORMAT = "2006-01-02 15:04:05"

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
	user.Password = utils.CryptoMd5(req.Password)
	user.CreateTime = now.Format(FORMAT)
	err = dao.DB.Create(user).Error
	return
}

// 用户登录

func UserLogin(req *model.UserLogin) (user *model.User, err error) {

	_, err = CheckUserExit(req.Email)
	if err != nil {
		return nil, err
	}
	password := utils.CryptoMd5(req.Password)
	result := dao.DB.Where("email = ? AND password = ? ", req.Email, password).Find(&user)
	fmt.Println(user.Username)
	return user, result.Error
}

// 通过id获取用户

func GetUserById(userId int64) (user model.User, err error) {
	result := dao.DB.Where("id = ?", userId).Find(&user)
	return user, result.Error
}

// 获取所有用户

func GetUserList() (users *[]model.User, err error) {
	tx := dao.DB.Find(&users)
	return users, tx.Error
}
