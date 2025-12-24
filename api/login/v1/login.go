// LoginReq 登录请求
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"登录" summary:"用户登录"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type LoginRes struct {
	Token   string `json:"token" dc:"JWT token for authentication"`
	Account string `json:"account"`
}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"注册" summary:"用户注册"`
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type RegisterRes struct {
	Account string `json:"account"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"登出" summary:"用户登出"`
}
type LogoutRes struct {
}

// 用邮箱登录（包含验证码）
type LoginByEmailReq struct {
	g.Meta `path:"/login/email" method:"post" tags:"登录" summary:"用户邮箱登录"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
	Code   string `json:"code" v:"required#验证码不能为空"`
}

type LoginByEmailRes struct {
	Token string `json:"token" dc:"JWT token for authentication"`
	Email string `json:"email"`
}

// 通过邮箱注册
type RegisterByEmailReq struct {
	g.Meta `path:"/register/email" method:"post" tags:"注册" summary:"用户邮箱注册"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
	Code   string `json:"code" v:"required#验证码不能为空"`
}

type RegisterByEmailRes struct {
	Email string `json:"email"`
}

// SendVerificationCodeReq 发送验证码
type SendVerificationCodeReq struct {
	g.Meta `path:"/email/send-code" method:"post" tags:"邮箱" summary:"发送验证码"`
	Email  string `json:"email" v:"email#邮箱格式不正确"`
}

type SendVerificationCodeRes struct {
	Success bool `json:"success"`
}
