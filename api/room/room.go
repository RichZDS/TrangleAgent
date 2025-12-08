// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package room

import (
	"context"

	"leke/api/room/v1"
)

type IRoomV1 interface {
	RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error)
	RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error)
	RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error)
	RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error)
	RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error)
}
