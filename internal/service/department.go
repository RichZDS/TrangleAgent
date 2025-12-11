// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/department/v1"
)

type (
	IDepartment interface {
		// DepartmentList 获取部门列表
		DepartmentList(ctx context.Context, req *v1.DepartmentListReq) (res *v1.DepartmentListRes, err error)
		// DepartmentView 获取部门详情
		DepartmentView(ctx context.Context, req *v1.DepartmentViewReq) (res *v1.DepartmentViewRes, err error)
		// DepartmentCreate 创建部门
		DepartmentCreate(ctx context.Context, req *v1.DepartmentCreateReq) (res *v1.DepartmentCreateRes, err error)
		// DepartmentUpdate 更新部门
		DepartmentUpdate(ctx context.Context, req *v1.DepartmentUpdateReq) (res *v1.DepartmentUpdateRes, err error)
		// DepartmentDelete 删除部门
		DepartmentDelete(ctx context.Context, req *v1.DepartmentDeleteReq) (res *v1.DepartmentDeleteRes, err error)
	}
)

var (
	localDepartment IDepartment
)

func Department() IDepartment {
	if localDepartment == nil {
		panic("implement not found for interface IDepartment, forgot register?")
	}
	return localDepartment
}

func RegisterDepartment(i IDepartment) {
	localDepartment = i
}
