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
