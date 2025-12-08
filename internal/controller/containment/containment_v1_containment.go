package containment

import (
	"context"
	"leke/api/containment"

	"leke/api/containment/v1"
	"leke/internal/service"
)

type ControllerV1 struct{}

func NewV1() containment.IContainmentV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error) {
	return service.Containment().ContainmentRepoList(ctx, req)
}
func (c *ControllerV1) ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error) {
	return service.Containment().ContainmentRepoView(ctx, req)
}
func (c *ControllerV1) ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error) {
	return service.Containment().ContainmentRepoUpdate(ctx, req)
}
func (c *ControllerV1) ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error) {
	return service.Containment().ContainmentRepoDelete(ctx, req)
}
