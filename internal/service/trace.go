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
	ITrace interface {
		// TraceList 获取轨迹列表
		TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error)
		// TraceView 查看轨迹详情
		TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error)
		// TraceUpdate 更新轨迹信息 (增加)
		TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error)
		// TraceReduce 减少轨迹数值
		TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error)
	}
)

var (
	localTrace ITrace
)

func Trace() ITrace {
	if localTrace == nil {
		panic("implement not found for interface ITrace, forgot register?")
	}
	return localTrace
}

func RegisterTrace(i ITrace) {
	localTrace = i
}
