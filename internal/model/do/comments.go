// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure of table comments for DAO operations like Where/Data.
type Comments struct {
	g.Meta     `orm:"table:comments, do:true"`
	Id         any         // 评论ID
	UserId     any         // 评论用户ID
	TargetType any         // 被评论对象类型（如 post/image/video 等）
	TargetId   any         // 被评论对象ID
	ParentId   any         // 父评论ID（回复某条评论时填写）
	RootId     any         // 根评论ID（同一楼的顶层评论ID，便于树查询）
	Content    any         // 评论内容
	Status     any         // 状态：1正常 0屏蔽 2删除
	LikeCount  any         // 点赞数
	ReplyCount any         // 回复数
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 软删除时间
}
