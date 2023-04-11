package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func RegisterHandler(c *gin.Context) {

	//var reg model.UserRegister
	userReg := new(model.UserRegister)
	err := c.ShouldBindJSON(&userReg)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler(c, "200", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}
	result, err := logic.CheckUserExit(userReg.Email)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println(result)
	if result.Id != 0 {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "当前用户已存在"})
		return
	}

	err = logic.CreateUser(userReg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功", "data": nil})
}

func LoginHandler(c *gin.Context) {
	var loginParams model.UserLogin
	err := c.ShouldBindJSON(&loginParams)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler(c, "200", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}
	err = logic.UserLogin(&loginParams)
	if err != nil {
		log.Printf("login登录失败%s", err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功"})
}
