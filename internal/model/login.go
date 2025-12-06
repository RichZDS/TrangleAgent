package model

type LoginField struct {
	Account  string `json:"account" v:"required#账号不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
	Nickname string `json:"nickname" v:"required#昵称不能为空"`
}
