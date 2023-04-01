package api

import (
	"awesome-go/m-oa/user/pkg/dao"
	"awesome-go/m-oa/user/pkg/model"
	"awesome-go/m-oa/user/pkg/repo"
	"common"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}

	mobile := ctx.PostForm("mobile")
	if common.VerifyMobile(mobile) {
		fmt.Println(mobile)
		ctx.JSON(http.StatusOK, rsp.Fail(model.NoLegalMobile, "手机号不合法"))
		return
	}

	code := "123456"
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功发送短信")

		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error("验证码存入redis出错:%v\n\n")
		}
	}()

	ctx.JSON(http.StatusOK, rsp.Success(code))
}
