// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComments is the golang structure of table forum_comments for DAO operations like Where/Data.
type ForumComments struct {
	g.Meta     `orm:"table:forum_comments, do:true"`
	Id         any         // 评论ID（主键）
	UserId     any         // 评论发布者ID（关联users.id，无外键）
	PostId     any         // 所属帖子ID（关联forum_posts.id，无外键）
	ParentId   any         // 父评论ID（NULL=根评论，非NULL=回复某条评论）
	Content    any         // 评论内容（支持emoji）
	Status     any         // 评论状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	LikeCount  any         // 点赞数
	ReplyCount any         // 回复数（冗余字段）
	CreatedAt  *gtime.Time // 评论创建时间
	UpdatedAt  *gtime.Time // 评论更新时间
	DeletedAt  *gtime.Time // 软删除时间
}
