package user

import (
	"context"

	"leke/api/user"
	v1 "leke/api/user/v1"
	"leke/internal/service"
)

type ControllerV1 struct{}

func NewV1() user.IUserV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	return service.User().UserList(ctx, req)
}
func (c *ControllerV1) UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error) {
	return service.User().UserView(ctx, req)
}
func (c *ControllerV1) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	return service.User().UserUpdate(ctx, req)
}
func (c *ControllerV1) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	return service.User().UserDelete(ctx, req)
}
