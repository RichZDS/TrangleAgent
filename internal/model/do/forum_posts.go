// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumPosts is the golang structure of table forum_posts for DAO operations like Where/Data.
type ForumPosts struct {
	g.Meta       `orm:"table:forum_posts, do:true"`
	Id           any         // 帖子ID（主键）
	UserId       any         // 发帖用户ID（关联users.id，无外键）
	Title        any         // 帖子标题
	Content      any         // 帖子正文（支持富文本/emoji）
	CoverImage   any         // 帖子封面图URL
	Status       any         // 帖子状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	ViewCount    any         // 浏览量（冗余字段）
	LikeCount    any         // 点赞数（冗余字段）
	CommentCount any         // 评论数（冗余字段）
	CreatedAt    *gtime.Time // 发帖时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 软删除时间
}
