// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GroupClassEnrollmentsDao is the data access object for the table group_class_enrollments.
type GroupClassEnrollmentsDao struct {
	table    string                       // table is the underlying table name of the DAO.
	group    string                       // group is the database configuration group name of the current DAO.
	columns  GroupClassEnrollmentsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler           // handlers for customized model modification.
}

// GroupClassEnrollmentsColumns defines and stores column names for the table group_class_enrollments.
type GroupClassEnrollmentsColumns struct {
	Id          string // 报名ID
	ClassId     string // 团课ID
	UserId      string // 用户ID
	Status      string // 状态：booked已报名，checked_in已签到，cancelled已取消
	EnrolledAt  string // 报名时间
	CheckedInAt string // 签到时间
}

// groupClassEnrollmentsColumns holds the columns for the table group_class_enrollments.
var groupClassEnrollmentsColumns = GroupClassEnrollmentsColumns{
	Id:          "id",
	ClassId:     "class_id",
	UserId:      "user_id",
	Status:      "status",
	EnrolledAt:  "enrolled_at",
	CheckedInAt: "checked_in_at",
}

// NewGroupClassEnrollmentsDao creates and returns a new DAO object for table data access.
func NewGroupClassEnrollmentsDao(handlers ...gdb.ModelHandler) *GroupClassEnrollmentsDao {
	return &GroupClassEnrollmentsDao{
		group:    "default",
		table:    "group_class_enrollments",
		columns:  groupClassEnrollmentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GroupClassEnrollmentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GroupClassEnrollmentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GroupClassEnrollmentsDao) Columns() GroupClassEnrollmentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GroupClassEnrollmentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GroupClassEnrollmentsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GroupClassEnrollmentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
