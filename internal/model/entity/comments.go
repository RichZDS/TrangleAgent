// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure for table comments.
type Comments struct {
	Id         uint64      `json:"id"         orm:"id"          description:"评论ID"`                          // 评论ID
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"评论用户ID"`                        // 评论用户ID
	TargetType string      `json:"targetType" orm:"target_type" description:"被评论对象类型（如 post/image/video 等）"` // 被评论对象类型（如 post/image/video 等）
	TargetId   uint64      `json:"targetId"   orm:"target_id"   description:"被评论对象ID"`                       // 被评论对象ID
	ParentId   uint64      `json:"parentId"   orm:"parent_id"   description:"父评论ID（回复某条评论时填写）"`              // 父评论ID（回复某条评论时填写）
	RootId     uint64      `json:"rootId"     orm:"root_id"     description:"根评论ID（同一楼的顶层评论ID，便于树查询）"`       // 根评论ID（同一楼的顶层评论ID，便于树查询）
	Content    string      `json:"content"    orm:"content"     description:"评论内容"`                          // 评论内容
	Status     int         `json:"status"     orm:"status"      description:"状态：1正常 0屏蔽 2删除"`                // 状态：1正常 0屏蔽 2删除
	LikeCount  uint        `json:"likeCount"  orm:"like_count"  description:"点赞数"`                           // 点赞数
	ReplyCount uint        `json:"replyCount" orm:"reply_count" description:"回复数"`                           // 回复数
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`                          // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`                          // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"软删除时间"`                         // 软删除时间
}
