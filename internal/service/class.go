// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/class/v1"
)

type (
	IClass interface {
		ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error)
		ClassView(ctx context.Context, req *v1.ClassViewReq) (res *v1.ClassViewRes, err error)
		ClassUpdate(ctx context.Context, req *v1.ClassUpdateReq) (res *v1.ClassUpdateRes, err error)
		ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error)
	}
)

var (
	localClass IClass
)

func Class() IClass {
	if localClass == nil {
		panic("implement not found for interface IClass, forgot register?")
	}
	return localClass
}

func RegisterClass(i IClass) {
	localClass = i
}
