package forum

import (
	"context"
	v1 "leke/api/forum/v1"
	"leke/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// ForumPostsCreate 创建帖子
func (c *ControllerV1) Create(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error) {
	return service.ForumPosts().Create(ctx, req)
}

// ForumPostsUpdate 更新帖子
func (c *ControllerV1) Update(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error) {
	return service.ForumPosts().Update(ctx, req)
}

// ForumPostsDelete 删除帖子
func (c *ControllerV1) Delete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error) {
	return service.ForumPosts().Delete(ctx, req)
}

// ForumPostsView 查看帖子详情
func (c *ControllerV1) View(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error) {
	return service.ForumPosts().View(ctx, req)
}

// ForumPostsList 获取帖子列表
func (c *ControllerV1) List(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error) {
	return service.ForumPosts().List(ctx, req)
}
