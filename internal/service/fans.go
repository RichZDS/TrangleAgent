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
	IFans interface {
		// FansList 获取粉丝列表
		//
		// 参数:
		//   - ctx context.Context: 上下文信息
		//   - req *v1.FansListReq: 粉丝列表请求参数
		//
		// 返回:
		//   - res *v1.FansListRes: 粉丝列表响应结果
		//   - err error: 错误信息
		FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error)
		// FansView 获取粉丝详情
		//
		// 参数:
		//   - ctx context.Context: 上下文信息
		//   - req *v1.FansViewReq: 粉丝详情请求参数
		//
		// 返回:
		//   - res *v1.FansViewRes: 粉丝详情响应结果
		//   - err error: 错误信息
		FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error)
		// FansUpdate 更新粉丝信息
		//
		// 参数:
		//   - ctx context.Context: 上下文信息
		//   - req *v1.FansUpdateReq: 粉丝更新请求参数
		//
		// 返回:
		//   - res *v1.FansUpdateRes: 粉丝更新响应结果
		//   - err error: 错误信息
		FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error)
		// FansDelete 删除粉丝关系
		//
		// 参数:
		//   - ctx context.Context: 上下文信息
		//   - req *v1.FansDeleteReq: 粉丝删除请求参数
		//
		// 返回:
		//   - res *v1.FansDeleteRes: 粉丝删除响应结果
		//   - err error: 错误信息
		FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error)
		// FansCreate 创建粉丝关系
		//
		// 参数:
		//   - ctx context.Context: 上下文信息
		//   - req *v1.FansCreateReq: 粉丝创建请求参数
		//
		// 返回:
		//   - res *v1.FansCreateRes: 粉丝创建响应结果
		//   - err error: 错误信息
		FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error)
	}
)

var (
	localFans IFans
)

func Fans() IFans {
	if localFans == nil {
		panic("implement not found for interface IFans, forgot register?")
	}
	return localFans
}

func RegisterFans(i IFans) {
	localFans = i
}
