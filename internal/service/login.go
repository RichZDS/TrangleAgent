// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/login/v1"
)

type (
	ILogin interface {
		Register(ctx context.Context, loginReq *v1.RegisterReq) (res *v1.RegisterRes, err error)
		Login(ctx context.Context, loginReq *v1.LoginReq) (res *v1.LoginRes, err error)
		// 通过邮箱注册
		RegisterByEmail(ctx context.Context, req *v1.RegisterByEmailReq) (res *v1.RegisterByEmailRes, err error)
		// 通过邮箱登录
		LoginByEmail(ctx context.Context, req *v1.LoginByEmailReq) (res *v1.LoginByEmailRes, err error)
		// 发送验证码
		SendVerificationCode(ctx context.Context, req *v1.SendVerificationCodeReq) (res *v1.SendVerificationCodeRes, err error)
		// 退出登录
		Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
