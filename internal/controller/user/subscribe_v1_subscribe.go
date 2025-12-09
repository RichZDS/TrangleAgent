package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "leke/api/user/v1"
)

// 注意：这个控制器已被弃用，请使用 user.ControllerV1 替代
type SubscribeControllerV1 struct{}

func NewSubscribeV1() *SubscribeControllerV1 {
	return &SubscribeControllerV1{}
}

func (c *SubscribeControllerV1) SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *SubscribeControllerV1) SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *SubscribeControllerV1) SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *SubscribeControllerV1) SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *SubscribeControllerV1) SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
