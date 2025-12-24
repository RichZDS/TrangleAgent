// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumPosts is the golang structure for table forum_posts.
type ForumPosts struct {
	Id           uint64      `json:"id"           orm:"id"            description:"帖子ID（主键）"`                                         // 帖子ID（主键）
	UserId       uint64      `json:"userId"       orm:"user_id"       description:"发帖用户ID（关联users.id，无外键）"`                           // 发帖用户ID（关联users.id，无外键）
	Title        string      `json:"title"        orm:"title"         description:"帖子标题"`                                             // 帖子标题
	Content      string      `json:"content"      orm:"content"       description:"帖子正文（支持富文本/emoji）"`                                // 帖子正文（支持富文本/emoji）
	CoverImage   string      `json:"coverImage"   orm:"cover_image"   description:"帖子封面图URL"`                                         // 帖子封面图URL
	Status       string      `json:"status"       orm:"status"        description:"帖子状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回"` // 帖子状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	ViewCount    uint        `json:"viewCount"    orm:"view_count"    description:"浏览量（冗余字段）"`                                        // 浏览量（冗余字段）
	LikeCount    uint        `json:"likeCount"    orm:"like_count"    description:"点赞数（冗余字段）"`                                        // 点赞数（冗余字段）
	CommentCount uint        `json:"commentCount" orm:"comment_count" description:"评论数（冗余字段）"`                                        // 评论数（冗余字段）
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"发帖时间"`                                             // 发帖时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`                                             // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"软删除时间"`                                            // 软删除时间
}
