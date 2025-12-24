// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComments is the golang structure for table forum_comments.
type ForumComments struct {
	Id         uint64      `json:"id"         orm:"id"          description:"评论ID（主键）"`                                         // 评论ID（主键）
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"评论发布者ID（关联users.id，无外键）"`                          // 评论发布者ID（关联users.id，无外键）
	PostId     uint64      `json:"postId"     orm:"post_id"     description:"所属帖子ID（关联forum_posts.id，无外键）"`                     // 所属帖子ID（关联forum_posts.id，无外键）
	ParentId   uint64      `json:"parentId"   orm:"parent_id"   description:"父评论ID（NULL=根评论，非NULL=回复某条评论）"`                     // 父评论ID（NULL=根评论，非NULL=回复某条评论）
	Content    string      `json:"content"    orm:"content"     description:"评论内容（支持emoji）"`                                    // 评论内容（支持emoji）
	Status     string      `json:"status"     orm:"status"      description:"评论状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回"` // 评论状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	LikeCount  uint        `json:"likeCount"  orm:"like_count"  description:"点赞数"`                                              // 点赞数
	ReplyCount uint        `json:"replyCount" orm:"reply_count" description:"回复数（冗余字段）"`                                        // 回复数（冗余字段）
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"评论创建时间"`                                           // 评论创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"评论更新时间"`                                           // 评论更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"软删除时间"`                                            // 软删除时间
}
