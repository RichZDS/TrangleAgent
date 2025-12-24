package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"leke/internal/model/response"
)

// ForumPostsCreateReq 创建帖子请求
type ForumPostsCreateReq struct {
	g.Meta     `path:"/forum/posts/create" method:"post" tags:"论坛帖子" summary:"创建帖子"`
	UserId     uint64 `json:"userId"      description:"发帖用户ID"`
	Title      string `json:"title"       description:"帖子标题"`
	Content    string `json:"content"     description:"帖子正文"`
	CoverImage string `json:"coverImage"  description:"帖子封面图URL"`
	Status     string `json:"status"      description:"帖子状态"`
}

// ForumPostsCreateRes 创建帖子响应
type ForumPostsCreateRes struct {
	Id uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsUpdateReq 更新帖子请求
type ForumPostsUpdateReq struct {
	g.Meta     `path:"/forum/posts/update" method:"put" tags:"论坛帖子" summary:"更新帖子"`
	Id         uint64 `json:"id"          description:"帖子ID"`
	UserId     uint64 `json:"userId,omitempty"      description:"发帖用户ID"`
	Title      string `json:"title,omitempty"       description:"帖子标题"`
	Content    string `json:"content,omitempty"     description:"帖子正文"`
	CoverImage string `json:"coverImage,omitempty"  description:"帖子封面图URL"`
	Status     string `json:"status,omitempty"      description:"帖子状态"`
}

// ForumPostsUpdateRes 更新帖子响应
type ForumPostsUpdateRes struct {
	Id uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsDeleteReq 删除帖子请求
type ForumPostsDeleteReq struct {
	g.Meta `path:"/forum/posts/delete" method:"delete" tags:"论坛帖子" summary:"删除帖子"`
	Id     uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsDeleteRes 删除帖子响应
type ForumPostsDeleteRes struct {
}

// ForumPostsViewReq 查看帖子请求
type ForumPostsViewReq struct {
	g.Meta `path:"/forum/posts/view" method:"get" tags:"论坛帖子" summary:"查看帖子"`
	Id     uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsViewRes 查看帖子响应
type ForumPostsViewRes struct {
	Id           uint64      `json:"id"           description:"帖子ID"`
	UserId       uint64      `json:"userId"       description:"发帖用户ID"`
	Title        string      `json:"title"        description:"帖子标题"`
	Content      string      `json:"content"      description:"帖子正文"`
	CoverImage   string      `json:"coverImage"   description:"帖子封面图URL"`
	Status       string      `json:"status"       description:"帖子状态"`
	ViewCount    uint        `json:"viewCount"    description:"浏览量"`
	LikeCount    uint        `json:"likeCount"    description:"点赞数"`
	CommentCount uint        `json:"commentCount" description:"评论数"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"发帖时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"更新时间"`
}

// ForumPostsListReq 帖子列表请求
type ForumPostsListReq struct {
	response.PageResult
	g.Meta `path:"/forum/posts/list" method:"get" tags:"论坛帖子" summary:"帖子列表"`
	UserId uint64 `json:"userId,omitempty" description:"发帖用户ID"`
	Title  string `json:"title,omitempty"  description:"帖子标题"`
	Status string `json:"status,omitempty" description:"帖子状态"`
}

// ForumPostsListRes 帖子列表响应
type ForumPostsListRes struct {
	response.PageResult
	List []*ForumPostsViewRes `json:"list" description:"帖子列表"`
}
