package user

import (
	"context"

	v1 "leke/api/user/v1"
	"leke/internal/service"
)

type RoleControllerV1 struct{}

func NewRoleV1() *RoleControllerV1 {
	return &RoleControllerV1{}
}

func (c *RoleControllerV1) RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	return service.User().RoleCreate(ctx, req)
}

func (c *RoleControllerV1) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	return service.User().RoleUpdate(ctx, req)
}

func (c *RoleControllerV1) RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error) {
	return service.User().RoleView(ctx, req)
}

func (c *RoleControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	return service.User().RoleList(ctx, req)
}

func (c *RoleControllerV1) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	return service.User().RoleDelete(ctx, req)
} // RolePermissionCheck 权限查询
func (c *RoleControllerV1) RolePermissionCheck(ctx context.Context, req *v1.RolePermissionCheckReq) (res *v1.RolePermissionCheckRes, err error) {
	return service.User().RolePermissionCheck(ctx, req)
}
