package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "leke/api/user/v1"
)

// 注意：这个控制器已被弃用，请使用 user.ControllerV1 替代
type FansControllerV1 struct{}

func NewFansV1() *FansControllerV1 {
	return &FansControllerV1{}
}

func (c *FansControllerV1) FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *FansControllerV1) FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *FansControllerV1) FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *FansControllerV1) FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *FansControllerV1) FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
