package controller

import (
	"bluebell/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(c *gin.Context) {

	var reg model.UserRegister
	err := c.ShouldBindJSON(reg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "请求参数有误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "请求成功"})
}
