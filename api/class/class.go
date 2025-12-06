// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package class

import (
	"context"

	"leke/api/class/v1"
)

type IClassV1 interface {
	ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error)
	ClassView(ctx context.Context, req *v1.ClassViewReq) (res *v1.ClassViewRes, err error)
	ClassUpdate(ctx context.Context, req *v1.ClassUpdateReq) (res *v1.ClassUpdateRes, err error)
	ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error)
}
