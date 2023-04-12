package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

func RegisterHandler(c *gin.Context) {
	userReg := new(model.UserRegister)
	err := c.ShouldBindJSON(&userReg)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}
	result, _ := logic.CheckUserExit(userReg.Email)
	if result.Id != 0 {
		ResponseError(c, CodeUserExits)
		return
	}

	err = logic.CreateUser(userReg)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}

func LoginHandler(c *gin.Context) {
	var loginParams model.UserLogin
	err := c.ShouldBindJSON(&loginParams)
	if err != nil {
		errs, _ := err.(validator.ValidationErrors)
		ResponseErrorWithMsg(c, CodeInvalidParams,
			utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		)
		return
	}
	err = logic.UserLogin(&loginParams)
	if err != nil {
		log.Printf("login登录失败%s", err)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}
