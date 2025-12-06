// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CoachesDao is the data access object for the table coaches.
type CoachesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CoachesColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CoachesColumns defines and stores column names for the table coaches.
type CoachesColumns struct {
	Id        string // 教练ID
	UserId    string // 对应的用户ID（如果教练也要登录的话）
	Name      string // 教练姓名
	Phone     string // 教练电话
	Specialty string // 擅长方向，比如瑜伽、力量、减脂
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// coachesColumns holds the columns for the table coaches.
var coachesColumns = CoachesColumns{
	Id:        "id",
	UserId:    "user_id",
	Name:      "name",
	Phone:     "phone",
	Specialty: "specialty",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCoachesDao creates and returns a new DAO object for table data access.
func NewCoachesDao(handlers ...gdb.ModelHandler) *CoachesDao {
	return &CoachesDao{
		group:    "default",
		table:    "coaches",
		columns:  coachesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CoachesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CoachesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CoachesDao) Columns() CoachesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CoachesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CoachesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CoachesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
