package v1

import (
	"leke/internal/model"
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// RoomListReq 房间列表请求
type RoomListReq struct {
	response.PageResult
	g.Meta     `path:"/room/list" method:"get" tags:"房间" summary:"房间列表"`
	RoomCode   string `json:"room_code"   orm:"room_code"    description:"房间号"`
	RoomName   string `json:"room_name"   orm:"room_name"    description:"房间名称"`
	HostId     uint64 `json:"host_id"     orm:"host_id"      description:"主持人用户ID"`
	Status     int    `json:"status"      orm:"status"       description:"房间状态：0未开始 1进行中 2已结束 3已关闭"`
	SystemName string `json:"system_name" orm:"system_name"  description:"规则系统，如COC、DND5E"`
}

// RoomListRes 房间列表响应
type RoomListRes struct {
	response.PageResult
	Rooms []*model.RoomParams `json:"rooms"`
}

// RoomViewReq 房间详情请求
type RoomViewReq struct {
	g.Meta   `path:"/room/view" method:"get" tags:"房间" summary:"房间详情"`
	Id       uint64 `json:"id"         orm:"id"          description:"房间ID"`
	RoomCode string `json:"room_code"  orm:"room_code"   description:"房间号"`
}

// RoomViewRes 房间详情响应
type RoomViewRes struct {
	model.RoomView
}

// RoomCreateReq 创建房间请求
type RoomCreateReq struct {
	g.Meta `path:"/room/create" method:"post" tags:"房间" summary:"创建房间"`

	// 这里根据你实际的 room 表字段自己增删
	RoomCode     string `json:"room_code"    orm:"room_code"     description:"房间号"`
	RoomName     string `json:"room_name"    orm:"room_name"     description:"房间名称"`
	HostId       uint64 `json:"host_id"      orm:"host_id"       description:"主持人用户ID"`
	MaxPlayers   uint   `json:"max_players"  orm:"max_players"   description:"最大玩家人数"`
	HasPassword  int    `json:"has_password" orm:"has_password"  description:"是否有密码：0无 1有"`
	RoomPassword string `json:"room_password" orm:"room_password" description:"房间密码(建议哈希)"`
	IsPrivate    int    `json:"is_private"   orm:"is_private"    description:"是否私密房：0公开 1私密"`
	SystemName   string `json:"system_name"  orm:"system_name"   description:"规则系统，如COC、DND5E"`
	ScenarioName string `json:"scenario_name" orm:"scenario_name" description:"模组/剧本名称"`
	Description  string `json:"description"  orm:"description"   description:"房间简介/招募说明"`
}

// RoomCreateRes 创建房间响应
type RoomCreateRes struct {
	Id       uint64 `json:"id"         orm:"id"          description:"房间ID"`
	RoomCode string `json:"room_code"  orm:"room_code"   description:"房间号"`
}

// RoomUpdateReq 更新房间请求
type RoomUpdateReq struct {
	g.Meta   `path:"/room/update" method:"put" tags:"房间" summary:"更新房间"`
	RoomCode string `json:"room_code"   orm:"room_code"   description:"房间号(作为更新定位字段)"`
	model.RoomParams
}

// RoomUpdateRes 更新房间响应
type RoomUpdateRes struct {
	Id uint64 `json:"id"         orm:"id"           description:"房间ID"`
}

// RoomDeleteReq 删除房间请求
type RoomDeleteReq struct {
	g.Meta   `path:"/room/delete" method:"delete" tags:"房间" summary:"删除房间"`
	RoomCode string `json:"room_code"   orm:"room_code"      description:"房间号"`
}

// RoomDeleteRes 删除房间响应
type RoomDeleteRes struct{}

type RoomJoinReq struct {
	g.Meta `path:"/room/join" method:"post" tags:"房间" summary:"加入房间"`
	RoomId uint64 `json:"roomId" description:"房间ID"`
}

type RoomJoinRes struct {
	Id       uint64 `json:"id"         description:"房间ID"`
	RoomCode string `json:"roomCode"   description:"房间号"`
}
