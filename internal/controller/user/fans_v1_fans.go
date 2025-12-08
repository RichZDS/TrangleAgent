package user

import (
	"context"

	"leke/api/user"
	v1 "leke/api/user/v1"
	"leke/internal/service"
)

type FansControllerV1 struct{}

func NewFansV1() user.IFansV1 {
	return &FansControllerV1{}
}

func (c *FansControllerV1) FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error) {
	return service.Fans().FansList(ctx, req)
}

func (c *FansControllerV1) FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error) {
	return service.Fans().FansView(ctx, req)
}

func (c *FansControllerV1) FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error) {
	return service.Fans().FansUpdate(ctx, req)
}

func (c *FansControllerV1) FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error) {
	return service.Fans().FansDelete(ctx, req)
}

func (c *FansControllerV1) FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error) {
	return service.Fans().FansCreate(ctx, req)
}
