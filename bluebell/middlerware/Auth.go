package middlerware

import (
	"bluebell/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

const CtxUserIdKey = "userId"
const CtxUserNameKey = "username"

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			utils.ResponseError(c, utils.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格进行分割
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.ResponseError(c, utils.CodeInvalidToken)
			c.Abort()
			return
		}
		token, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.ResponseError(c, utils.CodeInvalidToken)
			log.Printf("解析token失败%s", err)
			c.Abort()
			return
		}
		c.Set(CtxUserIdKey, token.UserId)
		c.Set(CtxUserNameKey, token.Username)
		c.Next()

	}
}
