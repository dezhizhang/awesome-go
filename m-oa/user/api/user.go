package api

import (
	"awesome-go/m-oa/user/pkg/dao"
	"awesome-go/m-oa/user/pkg/model"
	"awesome-go/m-oa/user/pkg/repo"
	"common"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Println("短信平台调用成功发送短信")
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Fatalf("验证码存入redis出错:%v\n\n", err)
		}
	}()

	ctx.JSON(http.StatusOK, rsp.Success(code))
}
