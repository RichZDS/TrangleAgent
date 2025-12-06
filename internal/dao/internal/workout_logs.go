// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WorkoutLogsDao is the data access object for the table workout_logs.
type WorkoutLogsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WorkoutLogsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WorkoutLogsColumns defines and stores column names for the table workout_logs.
type WorkoutLogsColumns struct {
	Id              string // 健身记录ID
	UserId          string // 关联的用户ID
	WorkoutDate     string // 健身日期
	StartTime       string // 开始时间
	EndTime         string // 结束时间
	DurationMinutes string // 时长（分钟）
	Calories        string // 估算消耗的卡路里
	Description     string // 本次训练的备注，比如训练内容、部位等
	CreatedAt       string // 记录创建时间
}

// workoutLogsColumns holds the columns for the table workout_logs.
var workoutLogsColumns = WorkoutLogsColumns{
	Id:              "id",
	UserId:          "user_id",
	WorkoutDate:     "workout_date",
	StartTime:       "start_time",
	EndTime:         "end_time",
	DurationMinutes: "duration_minutes",
	Calories:        "calories",
	Description:     "description",
	CreatedAt:       "created_at",
}

// NewWorkoutLogsDao creates and returns a new DAO object for table data access.
func NewWorkoutLogsDao(handlers ...gdb.ModelHandler) *WorkoutLogsDao {
	return &WorkoutLogsDao{
		group:    "default",
		table:    "workout_logs",
		columns:  workoutLogsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WorkoutLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WorkoutLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WorkoutLogsDao) Columns() WorkoutLogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WorkoutLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WorkoutLogsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WorkoutLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
