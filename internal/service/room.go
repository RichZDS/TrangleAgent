// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/room/v1"
)

type (
	IRoom interface {
		RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error)
		RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error)
		RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error)
		RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error)
		RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error)
		RoomJoin(ctx context.Context, req *v1.RoomJoinReq) (res *v1.RoomJoinRes, err error)
	}
)

var (
	localRoom IRoom
)

func Room() IRoom {
	if localRoom == nil {
		panic("implement not found for interface IRoom, forgot register?")
	}
	return localRoom
}

func RegisterRoom(i IRoom) {
	localRoom = i
}
