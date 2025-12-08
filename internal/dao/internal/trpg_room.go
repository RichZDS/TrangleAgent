// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrpgRoomDao is the data access object for the table trpg_room.
type TrpgRoomDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TrpgRoomColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TrpgRoomColumns defines and stores column names for the table trpg_room.
type TrpgRoomColumns struct {
	Id             string // 自增主键
	RoomId         string // 房间全局唯一ID（建议UUID）
	RoomCode       string // 房间号（玩家看到/输入的房间号）
	RoomName       string // 房间名称
	HostId         string // 主持人用户ID（GM/KP/经理，对应用户表id）
	HostNickname   string // 主持人昵称（冗余字段，可选）
	MaxPlayers     string // 最大玩家人数（不含主持人，可按需要约定）
	CurrentPlayers string // 当前玩家人数（不含主持人）
	HasPassword    string // 是否有密码：0无 1有
	RoomPassword   string // 房间密码（建议存加密/哈希后的密码）
	IsPrivate      string // 是否私密房：0公开 1私密
	Status         string // 房间状态：0未开始 1进行中 2已结束 3已关闭
	SystemName     string // 规则系统：如 COC、DND5E 等
	ScenarioName   string // 模组/剧本名称
	Description    string // 房间简介/招募说明
	CreatedAt      string // 房间创建时间
	StartedAt      string // 开团时间
	EndedAt        string // 结束时间
	UpdatedAt      string // 信息最近更新时间
}

// trpgRoomColumns holds the columns for the table trpg_room.
var trpgRoomColumns = TrpgRoomColumns{
	Id:             "id",
	RoomId:         "room_id",
	RoomCode:       "room_code",
	RoomName:       "room_name",
	HostId:         "host_id",
	HostNickname:   "host_nickname",
	MaxPlayers:     "max_players",
	CurrentPlayers: "current_players",
	HasPassword:    "has_password",
	RoomPassword:   "room_password",
	IsPrivate:      "is_private",
	Status:         "status",
	SystemName:     "system_name",
	ScenarioName:   "scenario_name",
	Description:    "description",
	CreatedAt:      "created_at",
	StartedAt:      "started_at",
	EndedAt:        "ended_at",
	UpdatedAt:      "updated_at",
}

// NewTrpgRoomDao creates and returns a new DAO object for table data access.
func NewTrpgRoomDao(handlers ...gdb.ModelHandler) *TrpgRoomDao {
	return &TrpgRoomDao{
		group:    "default",
		table:    "trpg_room",
		columns:  trpgRoomColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrpgRoomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrpgRoomDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrpgRoomDao) Columns() TrpgRoomColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrpgRoomDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrpgRoomDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TrpgRoomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
