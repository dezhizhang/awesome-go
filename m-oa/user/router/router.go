package router

import (
	"awesome-go/m-oa/user/api"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	rg := New()

	rg.Route(&api.RouterUser{}, r)
}

//func Register(ro ...Router) {
//	routers = append(routers, ro...)
//}
