package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handler(c *gin.Context, code string, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}
