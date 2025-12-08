package user

import (
	"context"

	"leke/api/user"
	v1 "leke/api/user/v1"
	"leke/internal/service"
)

type SubscribeControllerV1 struct{}

func NewSubscribeV1() user.ISubscribeV1 {
	return &SubscribeControllerV1{}
}

func (c *SubscribeControllerV1) SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error) {
	return service.Subscribe().SubscribeList(ctx, req)
}

func (c *SubscribeControllerV1) SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error) {
	return service.Subscribe().SubscribeView(ctx, req)
}

func (c *SubscribeControllerV1) SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error) {
	return service.Subscribe().SubscribeUpdate(ctx, req)
}

func (c *SubscribeControllerV1) SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error) {
	return service.Subscribe().SubscribeDelete(ctx, req)
}

func (c *SubscribeControllerV1) SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error) {
	return service.Subscribe().SubscribeCreate(ctx, req)
}
