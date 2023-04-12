package controller

import (
	"bluebell/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResCode int64

const (
	CodeSuccess = 100 + iota
	CodeInvalidParams
	CodeUserExits
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExits:       "用户已存在",
	CodeUserNotExit:     "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务器繁忙",
}

func getCode(code ResCode) string {
	return CodeMsgMap[code]
}

// 失败返回的数据

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &model.ResponseData{
		Code: int(code),
		Msg:  getCode(code),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &model.ResponseData{
		Code: int(code),
		Msg:  msg,
		Data: nil,
	})
}

// 成功后返回的数据

func ResponseSuccess(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &model.ResponseData{
		Code: int(code),
		Msg:  getCode(code),
		Data: data,
	})
}
