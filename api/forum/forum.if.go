// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package forum

import (
	"context"

	"leke/api/forum/v1"
)

type IForumV1 interface {
	ForumCommentsCreate(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error)
	ForumCommentsUpdate(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error)
	ForumCommentsDelete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error)
	ForumCommentsView(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error)
	ForumCommentsList(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error)
	ForumPostsCreate(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error)
	ForumPostsUpdate(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error)
	ForumPostsDelete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error)
	ForumPostsView(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error)
	ForumPostsList(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error)
}
