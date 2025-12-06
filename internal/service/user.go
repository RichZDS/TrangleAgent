// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/user/v1"
)

type (
	IUser interface {
		UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
		UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error)
		UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error)
		UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
