// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/forum/v1"
)

type (
	IForumComments interface {
		// ForumCommentsCreate 创建评论
		Create(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error)
		// ForumCommentsUpdate 更新评论
		Update(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error)
		// ForumCommentsView 查看评论
		View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error)
		// ForumCommentsList 评论列表
		List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error)
		// ForumCommentsDelete 删除评论
		Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error)
	}
	IForumPosts interface {
		// ForumPostsCreate 创建帖子
		Create(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error)
		// ForumPostsUpdate 更新帖子
		Update(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error)
		// ForumPostsView 查看帖子
		View(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error)
		// ForumPostsList 帖子列表
		List(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error)
		// ForumPostsDelete 删除帖子
		Delete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error)
	}
)

var (
	localForumComments IForumComments
	localForumPosts    IForumPosts
)

func ForumComments() IForumComments {
	if localForumComments == nil {
		panic("implement not found for interface IForumComments, forgot register?")
	}
	return localForumComments
}

func RegisterForumComments(i IForumComments) {
	localForumComments = i
}

func ForumPosts() IForumPosts {
	if localForumPosts == nil {
		panic("implement not found for interface IForumPosts, forgot register?")
	}
	return localForumPosts
}

func RegisterForumPosts(i IForumPosts) {
	localForumPosts = i
}
