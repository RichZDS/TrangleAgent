// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoom is the golang structure for table trpg_room.
type TrpgRoom struct {
	Id             uint64      `json:"id"             orm:"id"              description:"自增主键"`                      // 自增主键
	RoomId         string      `json:"roomId"         orm:"room_id"         description:"房间全局唯一ID（建议UUID）"`          // 房间全局唯一ID（建议UUID）
	RoomCode       string      `json:"roomCode"       orm:"room_code"       description:"房间号（玩家看到/输入的房间号）"`          // 房间号（玩家看到/输入的房间号）
	RoomName       string      `json:"roomName"       orm:"room_name"       description:"房间名称"`                      // 房间名称
	HostId         uint64      `json:"hostId"         orm:"host_id"         description:"主持人用户ID（GM/KP/经理，对应用户表id）"` // 主持人用户ID（GM/KP/经理，对应用户表id）
	HostNickname   string      `json:"hostNickname"   orm:"host_nickname"   description:"主持人昵称（冗余字段，可选）"`            // 主持人昵称（冗余字段，可选）
	MaxPlayers     uint        `json:"maxPlayers"     orm:"max_players"     description:"最大玩家人数（不含主持人，可按需要约定）"`      // 最大玩家人数（不含主持人，可按需要约定）
	CurrentPlayers uint        `json:"currentPlayers" orm:"current_players" description:"当前玩家人数（不含主持人）"`             // 当前玩家人数（不含主持人）
	HasPassword    int         `json:"hasPassword"    orm:"has_password"    description:"是否有密码：0无 1有"`               // 是否有密码：0无 1有
	RoomPassword   string      `json:"roomPassword"   orm:"room_password"   description:"房间密码（建议存加密/哈希后的密码）"`        // 房间密码（建议存加密/哈希后的密码）
	IsPrivate      int         `json:"isPrivate"      orm:"is_private"      description:"是否私密房：0公开 1私密"`             // 是否私密房：0公开 1私密
	Status         int         `json:"status"         orm:"status"          description:"房间状态：0未开始 1进行中 2已结束 3已关闭"`  // 房间状态：0未开始 1进行中 2已结束 3已关闭
	SystemName     string      `json:"systemName"     orm:"system_name"     description:"规则系统：如 COC、DND5E 等"`        // 规则系统：如 COC、DND5E 等
	ScenarioName   string      `json:"scenarioName"   orm:"scenario_name"   description:"模组/剧本名称"`                   // 模组/剧本名称
	Description    string      `json:"description"    orm:"description"     description:"房间简介/招募说明"`                 // 房间简介/招募说明
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"房间创建时间"`                    // 房间创建时间
	StartedAt      *gtime.Time `json:"startedAt"      orm:"started_at"      description:"开团时间"`                      // 开团时间
	EndedAt        *gtime.Time `json:"endedAt"        orm:"ended_at"        description:"结束时间"`                      // 结束时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"信息最近更新时间"`                  // 信息最近更新时间
}
