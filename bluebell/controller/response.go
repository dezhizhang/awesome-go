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

func FailHandler(c *gin.Context, code int) {
	rsp := &model.ResponseData{
		Code: code,
		Msg:  getCode(code),
		Data: nil,
	}
	c.JSON(http.StatusOK, rsp)
}
