// LoginReq 登录请求
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"Login" summary:"User Login"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	Token   string `json:"token" dc:"JWT token for authentication"`
	Account string `json:"account"`
}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"Register" summary:"User Register"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type RegisterRes struct {
	Account string `json:"account"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"Logout" summary:"User Logout"`
}
type LogoutRes struct {
}
