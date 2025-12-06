package class

import (
	"context"
	"leke/api/class"
	service "leke/internal/service"

	"leke/api/class/v1"
)

type ControllerV1 struct{}

func NewV1() class.IClassV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error) {
	res, err = service.Class().ClassList(ctx, req)
	return
}
func (c *ControllerV1) ClassView(ctx context.Context, req *v1.ClassViewReq) (res *v1.ClassViewRes, err error) {
	res, err = service.Class().ClassView(ctx, req)
	return
}
func (c *ControllerV1) ClassUpdate(ctx context.Context, req *v1.ClassUpdateReq) (res *v1.ClassUpdateRes, err error) {
	res, err = service.Class().ClassUpdate(ctx, req)
	return
}
func (c *ControllerV1) ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error) {
	res, err = service.Class().ClassDelete(ctx, req)
	return
}
