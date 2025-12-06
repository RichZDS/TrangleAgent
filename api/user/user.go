// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"leke/api/user/v1"
)

type IUserV1 interface {
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error)
	UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error)
	UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error)
}
