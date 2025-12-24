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
		// RoleCreate 创建角色
		RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error)
		// RoleUpdate 更新角色
		RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error)
		// RoleView 查看角色详情
		RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error)
		// RoleList 获取角色列表
		RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
		// RoleDelete 删除角色
		RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error)
		// RolePermissionCheck 权限查询
		RolePermissionCheck(ctx context.Context, req *v1.RolePermissionCheckReq) (res *v1.RolePermissionCheckRes, err error)
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
