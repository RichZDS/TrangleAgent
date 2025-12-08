// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package containment

import (
	"context"

	"leke/api/containment/v1"
)

type IContainmentV1 interface {
	ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error)
	ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error)
	ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error)
	ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error)
}
