package controller

import (
	"bluebell/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess = 100 + iota
	CodeInvalidParams
)

func getCode(code int) string {
	return ""
}

// 失败返回的数据

func FailHandler(c *gin.Context, code int) {
	resp := &model.ResponseData{
		Code: code,
		Msg:  getCode(code),
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

// 成功后返回的数据

func SuccessHandler(c *gin.Context, code int, data interface{}) {
	resp := &model.ResponseData{
		Code: code,
		Msg:  getCode(code),
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}
