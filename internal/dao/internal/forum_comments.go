// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumCommentsDao is the data access object for the table forum_comments.
type ForumCommentsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ForumCommentsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ForumCommentsColumns defines and stores column names for the table forum_comments.
type ForumCommentsColumns struct {
	Id         string // 评论ID（主键）
	UserId     string // 评论发布者ID（关联users.id，无外键）
	PostId     string // 所属帖子ID（关联forum_posts.id，无外键）
	ParentId   string // 父评论ID（NULL=根评论，非NULL=回复某条评论）
	Content    string // 评论内容（支持emoji）
	Status     string // 评论状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	LikeCount  string // 点赞数
	ReplyCount string // 回复数（冗余字段）
	CreatedAt  string // 评论创建时间
	UpdatedAt  string // 评论更新时间
	DeletedAt  string // 软删除时间
}

// forumCommentsColumns holds the columns for the table forum_comments.
var forumCommentsColumns = ForumCommentsColumns{
	Id:         "id",
	UserId:     "user_id",
	PostId:     "post_id",
	ParentId:   "parent_id",
	Content:    "content",
	Status:     "status",
	LikeCount:  "like_count",
	ReplyCount: "reply_count",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewForumCommentsDao creates and returns a new DAO object for table data access.
func NewForumCommentsDao(handlers ...gdb.ModelHandler) *ForumCommentsDao {
	return &ForumCommentsDao{
		group:    "default",
		table:    "forum_comments",
		columns:  forumCommentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumCommentsDao) Columns() ForumCommentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ForumCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
