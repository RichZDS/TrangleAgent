package room

import (
	"context"
	"leke/api/room"
	v1 "leke/api/room/v1"
	"leke/internal/service"
)

type ControllerV1 struct{}

func NewV1() room.IRoomV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error) {
	return service.Room().RoomList(ctx, req)
}
func (c *ControllerV1) RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error) {
	return service.Room().RoomView(ctx, req)
}

func (c *ControllerV1) RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error) {
	return service.Room().RoomUpdate(ctx, req)
}
func (c *ControllerV1) RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error) {
	return service.Room().RoomDelete(ctx, req)
}
func (c *ControllerV1) RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error) {
	return service.Room().RoomCreate(ctx, req)
}

func (c *ControllerV1) RoomJoin(ctx context.Context, req *v1.RoomJoinReq) (res *v1.RoomJoinRes, err error) {
	return service.Room().RoomJoin(ctx, req)
}
