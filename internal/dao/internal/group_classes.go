// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GroupClassesDao is the data access object for the table group_classes.
type GroupClassesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  GroupClassesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// GroupClassesColumns defines and stores column names for the table group_classes.
type GroupClassesColumns struct {
	Id          string // 团课ID
	Title       string // 课程名称，如：燃脂搏击、瑜伽
	Description string // 课程介绍
	CoachId     string // 授课教练ID
	Location    string // 上课地点/门店/教室
	StartTime   string // 开始时间
	EndTime     string // 结束时间
	MaxCapacity string // 最大人数
	Price       string // 价格（0表示免费或已包含在会员卡内）
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// groupClassesColumns holds the columns for the table group_classes.
var groupClassesColumns = GroupClassesColumns{
	Id:          "id",
	Title:       "title",
	Description: "description",
	CoachId:     "coach_id",
	Location:    "location",
	StartTime:   "start_time",
	EndTime:     "end_time",
	MaxCapacity: "max_capacity",
	Price:       "price",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewGroupClassesDao creates and returns a new DAO object for table data access.
func NewGroupClassesDao(handlers ...gdb.ModelHandler) *GroupClassesDao {
	return &GroupClassesDao{
		group:    "default",
		table:    "group_classes",
		columns:  groupClassesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GroupClassesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GroupClassesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GroupClassesDao) Columns() GroupClassesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GroupClassesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GroupClassesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GroupClassesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
