package controller

import (
	"leke/internal/controller/class"
	"leke/internal/controller/login"
	"leke/internal/controller/user"

	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterControllers 将所有控制器绑定到路由组
func RegisterControllers(group *ghttp.RouterGroup) {
	group.Bind(
		login.NewV1(),
		class.NewV1(),
		user.NewV1(),
	)
}
