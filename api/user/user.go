// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"leke/api/user/v1"
)

type IUserV1 interface {
	FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error)
	FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error)
	FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error)
	FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error)
	FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error)
	SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error)
	SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error)
	SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error)
	SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error)
	SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error)
	TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error)
	TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error)
	TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error)
	TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error)
	UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error)
	UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error)
}
