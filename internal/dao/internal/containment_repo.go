// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContainmentRepoDao is the data access object for the table containment_repo.
type ContainmentRepoDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  ContainmentRepoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// ContainmentRepoColumns defines and stores column names for the table containment_repo.
type ContainmentRepoColumns struct {
	Id          string // auto-increment primary key
	TerminalId  string // terminal (散逸端)
	Abnormal    string // number of contained anomalies (收容异常的数量)
	AnomalyName string // name of the anomaly (异常体的名字)
	AgentName   string // agent (特工)
	RepoName    string // containment repository name or code (收容库)
}

// containmentRepoColumns holds the columns for the table containment_repo.
var containmentRepoColumns = ContainmentRepoColumns{
	Id:          "id",
	TerminalId:  "terminal_id",
	Abnormal:    "abnormal",
	AnomalyName: "anomaly_name",
	AgentName:   "agent_name",
	RepoName:    "repo_name",
}

// NewContainmentRepoDao creates and returns a new DAO object for table data access.
func NewContainmentRepoDao(handlers ...gdb.ModelHandler) *ContainmentRepoDao {
	return &ContainmentRepoDao{
		group:    "default",
		table:    "containment_repo",
		columns:  containmentRepoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContainmentRepoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContainmentRepoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContainmentRepoDao) Columns() ContainmentRepoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContainmentRepoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContainmentRepoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ContainmentRepoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
