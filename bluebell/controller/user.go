package controller

import (
	"bluebell/logic"
	"bluebell/middlerware"
	"bluebell/model"
	"bluebell/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

// 用户注册

func RegisterHandler(c *gin.Context) {
	userReg := new(model.UserRegister)
	err := c.ShouldBindJSON(&userReg)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, utils.CodeInvalidParams)
			return
		}
		utils.ResponseErrorWithMsg(c, utils.CodeInvalidParams,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}
	result, _ := logic.CheckUserExit(userReg.Email)
	if result.Id != 0 {
		utils.ResponseError(c, utils.CodeUserExits)
		return
	}

	err = logic.CreateUser(userReg)
	if err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess, nil)
}

// 用户登录

func LoginHandler(c *gin.Context) {
	var loginParams model.UserLogin
	err := c.ShouldBindJSON(&loginParams)
	if err != nil {
		errs, _ := err.(validator.ValidationErrors)
		utils.ResponseErrorWithMsg(c, utils.CodeInvalidParams,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}
	user, err := logic.UserLogin(&loginParams)
	if err != nil && user.Username == "" {
		log.Printf("login登录失败%s", err)
		return
	}
	token, err := utils.GenToken(user.Id, user.Username)
	if err != nil {
		log.Printf("生成token失败%s", err)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess, token)
}

// 通过id获取用户信息

func UserByIdHandler(c *gin.Context) {
	value, ok := c.Get(middlerware.CtxUserIdKey)
	if !ok {
		log.Printf("中间件获取id失败%s", ok)
		return
	}
	userId, ok := value.(int64)
	if !ok {
		log.Printf("用户id转换失败")
		return
	}
	user, err := logic.GetUserById(userId)
	if err != nil {
		log.Printf("通过UserByIdHandler获取用户失败%s", err)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess, user)
}

// 获取用户列表

func UserListHandler(c *gin.Context) {
	fmt.Println(c.Get(middlerware.CtxUserIdKey))
}
