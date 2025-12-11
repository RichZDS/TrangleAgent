package containment

import (
	"context"
	"leke/api/containment"

	"leke/api/containment/v1"
	"leke/internal/service"
)

// ControllerV1 控制器结构体
type ControllerV1 struct{}

// NewV1 创建一个新的v1版本控制器实例
//
// 返回:
//   - containment.IContainmentV1: v1版本控制器接口实现
func NewV1() containment.IContainmentV1 {
	return &ControllerV1{}
}

// ContainmentRepoList 获取收容库列表
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.ContainmentRepoListReq: 收容库列表请求参数
//
// 返回:
//   - res *v1.ContainmentRepoListRes: 收容库列表响应结果
//   - err error: 错误信息
func (c *ControllerV1) ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error) {
	return service.Containment().ContainmentRepoList(ctx, req)
}

// ContainmentRepoView 查看收容库详情
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.ContainmentRepoViewReq: 收容库详情请求参数
//
// 返回:
//   - res *v1.ContainmentRepoViewRes: 收容库详情响应结果
//   - err error: 错误信息
func (c *ControllerV1) ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error) {
	return service.Containment().ContainmentRepoView(ctx, req)
}

// ContainmentRepoUpdate 更新收容库信息
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.ContainmentRepoUpdateReq: 收容库更新请求参数
//
// 返回:
//   - res *v1.ContainmentRepoUpdateRes: 收容库更新响应结果
//   - err error: 错误信息
func (c *ControllerV1) ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error) {
	return service.Containment().ContainmentRepoUpdate(ctx, req)
}

// ContainmentRepoDelete 删除收容库
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.ContainmentRepoDeleteReq: 收容库删除请求参数
//
// 返回:
//   - res *v1.ContainmentRepoDeleteRes: 收容库删除响应结果
//   - err error: 错误信息
func (c *ControllerV1) ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error) {
	return service.Containment().ContainmentRepoDelete(ctx, req)
}
