package forum

import (
	"context"
	v1 "leke/api/forum/v1"
	"leke/internal/service"
)

// ForumCommentsCreate 创建评论
func (c *ControllerV1) ForumCommentsCreate(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error) {
	return service.ForumComments().Create(ctx, req)
}

// ForumCommentsUpdate 更新评论
func (c *ControllerV1) ForumCommentsUpdate(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error) {
	return service.ForumComments().Update(ctx, req)
}

// ForumCommentsDelete 删除评论
func (c *ControllerV1) ForumCommentsDelete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error) {
	return service.ForumComments().Delete(ctx, req)
}

// ForumCommentsView 查看评论详情
func (c *ControllerV1) ForumCommentsView(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error) {
	return service.ForumComments().View(ctx, req)
}

// ForumCommentsList 获取评论列表
func (c *ControllerV1) ForumCommentsList(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error) {
	return service.ForumComments().List(ctx, req)
}
