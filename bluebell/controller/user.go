package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func RegisterHandler(c *gin.Context) {

	//var reg model.UserRegister
	user := new(model.UserRegister)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
		return
	}
	id, err := logic.CheckUserExit(user.Email)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	if id != 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "当前用户已存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "请求成功"})
}
