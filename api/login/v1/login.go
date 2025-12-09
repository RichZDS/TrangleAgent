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

// 用邮箱登录（包含验证码）
type LoginByEmailReq struct {
	g.Meta `path:"/login/email" method:"post" tags:"Login" summary:"User Login By Email"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
	Code   string `json:"code" v:"required#验证码不能为空"`
}

type LoginByEmailRes struct {
	Token string `json:"token" dc:"JWT token for authentication"`
	Email string `json:"email"`
}

// 通过邮箱注册
type RegisterByEmailReq struct {
	g.Meta `path:"/register/email" method:"post" tags:"Register" summary:"User Register By Email"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
	Code   string `json:"code" v:"required#验证码不能为空"`
}

type RegisterByEmailRes struct {
	Email string `json:"email"`
}

// SendVerificationCodeReq 发送验证码
type SendVerificationCodeReq struct {
	g.Meta `path:"/email/send-code" method:"post" tags:"Email" summary:"Send Verification Code"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
}

type SendVerificationCodeRes struct {
	Success bool `json:"success"`
}
