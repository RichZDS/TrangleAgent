package user

import (
	"context"

	v1 "leke/api/user/v1"
	"leke/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
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

// Fans 相关方法
func (c *ControllerV1) FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error) {
	return service.Fans().FansList(ctx, req)
}

func (c *ControllerV1) FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error) {
	return service.Fans().FansView(ctx, req)
}

func (c *ControllerV1) FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error) {
	return service.Fans().FansUpdate(ctx, req)
}

func (c *ControllerV1) FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error) {
	return service.Fans().FansDelete(ctx, req)
}

func (c *ControllerV1) FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error) {
	return service.Fans().FansCreate(ctx, req)
}

// Subscribe 相关方法
func (c *ControllerV1) SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error) {
	return service.Subscribe().SubscribeList(ctx, req)
}

func (c *ControllerV1) SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error) {
	return service.Subscribe().SubscribeView(ctx, req)
}

func (c *ControllerV1) SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error) {
	return service.Subscribe().SubscribeUpdate(ctx, req)
}

func (c *ControllerV1) SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error) {
	return service.Subscribe().SubscribeDelete(ctx, req)
}

func (c *ControllerV1) SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error) {
	return service.Subscribe().SubscribeCreate(ctx, req)
}

// Trace 相关方法
func (c *ControllerV1) TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error) {
	return service.Trace().TraceList(ctx, req)
}

func (c *ControllerV1) TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error) {
	return service.Trace().TraceView(ctx, req)
}

func (c *ControllerV1) TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error) {
	return service.Trace().TraceUpdate(ctx, req)
}

func (c *ControllerV1) TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error) {
	return service.Trace().TraceReduce(ctx, req)
}
