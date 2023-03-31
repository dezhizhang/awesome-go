package api

import (
	"common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}

	//mobile := ctx.PostForm("mobile")
	//if !common.VerifyMobile(mobile) {
	//	ctx.JSON(http.StatusOK, rsp.Fail(common.NoLegalMobile, "手机号不合法"))
	//	return
	//}

	code := "123456"
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("短信平台调用成功发送短信")
	}()

	ctx.JSON(http.StatusOK, rsp.Success(code))
}
