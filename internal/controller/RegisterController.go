package controller

import (
	"leke/internal/controller/containment"
	"leke/internal/controller/department"
	"leke/internal/controller/login"
	"leke/internal/controller/room"
	"leke/internal/controller/user"
	"leke/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterControllers 将所有控制器绑定到路由组
func RegisterControllers(group *ghttp.RouterGroup) {
	// 登录相关接口不需要JWT验证
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(
			login.NewV1(),
		)
	})

	// 其他需要JWT验证的接口
	group.Group("/", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.JWTAuth)
		g.Bind(
			user.NewV1(),
			user.NewRoleV1(),
			department.NewV1(),
			containment.NewV1(),
			room.NewV1(),
		)
	})
}
