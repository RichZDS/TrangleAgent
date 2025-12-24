// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumPostsDao is the data access object for the table forum_posts.
type ForumPostsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ForumPostsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ForumPostsColumns defines and stores column names for the table forum_posts.
type ForumPostsColumns struct {
	Id           string // 帖子ID（主键）
	UserId       string // 发帖用户ID（关联users.id，无外键）
	Title        string // 帖子标题
	Content      string // 帖子正文（支持富文本/emoji）
	CoverImage   string // 帖子封面图URL
	Status       string // 帖子状态：normal=正常 deleted=软删除 audit=审核中 reject=审核驳回
	ViewCount    string // 浏览量（冗余字段）
	LikeCount    string // 点赞数（冗余字段）
	CommentCount string // 评论数（冗余字段）
	CreatedAt    string // 发帖时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 软删除时间
}

// forumPostsColumns holds the columns for the table forum_posts.
var forumPostsColumns = ForumPostsColumns{
	Id:           "id",
	UserId:       "user_id",
	Title:        "title",
	Content:      "content",
	CoverImage:   "cover_image",
	Status:       "status",
	ViewCount:    "view_count",
	LikeCount:    "like_count",
	CommentCount: "comment_count",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewForumPostsDao creates and returns a new DAO object for table data access.
func NewForumPostsDao(handlers ...gdb.ModelHandler) *ForumPostsDao {
	return &ForumPostsDao{
		group:    "default",
		table:    "forum_posts",
		columns:  forumPostsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumPostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumPostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumPostsDao) Columns() ForumPostsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumPostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumPostsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumPostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
