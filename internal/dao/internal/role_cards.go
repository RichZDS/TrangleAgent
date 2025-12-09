// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleCardsDao is the data access object for the table role_cards.
type RoleCardsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RoleCardsColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RoleCardsColumns defines and stores column names for the table role_cards.
type RoleCardsColumns struct {
	Id             string // 角色卡ID
	UserId         string // 所属用户ID
	DepartmentId   string // 所属部门ID
	Commendation   string // 嘉奖次数
	Reprimand      string // 申戒次数
	BlueTrack      string // 蓝轨（0-40）
	YellowTrack    string // 黄轨（0-40）
	RedTrack       string // 红轨（0-40）
	ArcAbnormal    string // ARC：异常
	ArcReality     string // ARC：现实
	ArcPosition    string // ARC：职位
	AgentName      string // 特工名字
	CodeName       string // 代号
	Gender         string // 性别
	QaMeticulous   string // Meticulousness (0-100, QA)
	QaDeception    string // Deception (0-100, QA)
	QaVigor        string // Vigor / Drive (0-100, QA)
	QaEmpathy      string // Empathy (0-100, QA)
	QaInitiative   string // Initiative (0-100, QA)
	QaResilience   string // Resilience / Persistence (0-100, QA)
	QaPresence     string // Presence / Charisma (0-100, QA)
	QaProfessional string // Professionalism (0-100, QA)
	QaDiscretion   string // Discretion / Low profile (0-100, QA)
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// roleCardsColumns holds the columns for the table role_cards.
var roleCardsColumns = RoleCardsColumns{
	Id:             "id",
	UserId:         "user_id",
	DepartmentId:   "department_id",
	Commendation:   "commendation",
	Reprimand:      "reprimand",
	BlueTrack:      "blue_track",
	YellowTrack:    "yellow_track",
	RedTrack:       "red_track",
	ArcAbnormal:    "arc_abnormal",
	ArcReality:     "arc_reality",
	ArcPosition:    "arc_position",
	AgentName:      "agent_name",
	CodeName:       "code_name",
	Gender:         "gender",
	QaMeticulous:   "qa_meticulous",
	QaDeception:    "qa_deception",
	QaVigor:        "qa_vigor",
	QaEmpathy:      "qa_empathy",
	QaInitiative:   "qa_initiative",
	QaResilience:   "qa_resilience",
	QaPresence:     "qa_presence",
	QaProfessional: "qa_professional",
	QaDiscretion:   "qa_discretion",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewRoleCardsDao creates and returns a new DAO object for table data access.
func NewRoleCardsDao(handlers ...gdb.ModelHandler) *RoleCardsDao {
	return &RoleCardsDao{
		group:    "default",
		table:    "role_cards",
		columns:  roleCardsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoleCardsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoleCardsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoleCardsDao) Columns() RoleCardsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoleCardsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoleCardsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RoleCardsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
