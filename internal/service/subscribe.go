// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/user/v1"
)

type (
	ISubscribe interface {
		// SubscribeList 获取关注列表
		SubscribeList(ctx context.Context, req *v1.SubscribeListReq) (res *v1.SubscribeListRes, err error)
		// SubscribeView 获取关注详情
		SubscribeView(ctx context.Context, req *v1.SubscribeViewReq) (res *v1.SubscribeViewRes, err error)
		// SubscribeUpdate 更新关注信息
		SubscribeUpdate(ctx context.Context, req *v1.SubscribeUpdateReq) (res *v1.SubscribeUpdateRes, err error)
		// SubscribeDelete 删除关注关系
		SubscribeDelete(ctx context.Context, req *v1.SubscribeDeleteReq) (res *v1.SubscribeDeleteRes, err error)
		// SubscribeCreate 创建关注关系
		SubscribeCreate(ctx context.Context, req *v1.SubscribeCreateReq) (res *v1.SubscribeCreateRes, err error)
	}
)

var (
	localSubscribe ISubscribe
)

func Subscribe() ISubscribe {
	if localSubscribe == nil {
		panic("implement not found for interface ISubscribe, forgot register?")
	}
	return localSubscribe
}

func RegisterSubscribe(i ISubscribe) {
	localSubscribe = i
}
