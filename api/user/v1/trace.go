package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"leke/internal/model/response"
)

// TraceListReq 查询轨迹列表请求参数
type TraceListReq struct {
	response.PageResult
	g.Meta       `path:"/trace/list" method:"get" tags:"轨迹" summary:"轨迹列表"`
	UserId       uint64 `json:"userId"       description:"用户ID"`
	RoleId       uint64 `json:"roleId"       description:"角色ID"`
	DepartmentId uint64 `json:"departmentId"     description:"部门ID"`
	RedTrace     int    `json:"redTrace"     description:"红轨"`
	YellowTrace  int    `json:"yellowTrace"     description:"黄轨"`
	BlueTrace    int    `json:"blueTrace"     description:"蓝轨"`
}

// TraceListRes 查询轨迹列表响应数据
type TraceListRes struct {
	response.PageResult
	List []*TraceViewParams `json:"list"`
}

// TraceViewParams 轨迹视图参数
type TraceViewParams struct {
	UserId       uint64 `json:"userId"       description:"用户ID"`
	RoleId       uint64 `json:"roleId"       description:"角色ID"`
	DepartmentId uint64 `json:"departmentId"     description:"部门ID"`
	RedTrace     int    `json:"redTrace"     description:"红轨"`
	YellowTrace  int    `json:"yellowTrace"     description:"黄轨"`
	BlueTrace    int    `json:"blueTrace"     description:"蓝轨"`
}

// TraceViewReq 查看轨迹详情请求参数
type TraceViewReq struct {
	g.Meta      `path:"/trace/view" method:"get" tags:"轨迹" summary:"轨迹详情"`
	RoleId      uint64 `json:"roleId"       description:"角色ID"`
	RedTrace    int    `json:"redTrace" description:"红轨"`
	YellowTrace int    `json:"yellowTrace" description:"黄轨"`
	BlueTrace   int    `json:"blueTrace" description:"蓝轨"`
}

// TraceViewRes 查看轨迹详情响应数据
type TraceViewRes struct {
	TraceViewParams
}

// TraceUpdateReq 更新轨迹请求参数
type TraceUpdateReq struct {
	g.Meta       `path:"/trace/update" method:"put" tags:"轨迹" summary:"更新轨迹"`
	UserId       uint64 `json:"userId"`
	RoleId       uint64 `json:"roleId"       description:"角色ID"`
	DepartmentId uint64 `json:"departmentId"`
	RedTrace     int    `json:"redTrace" description:"红轨"`
	YellowTrace  int    `json:"yellowTrace" description:"黄轨"`
	BlueTrace    int    `json:"blueTrace" description:"蓝轨"`
}

// TraceUpdateRes 更新轨迹响应数据
type TraceUpdateRes struct {
	RoleId      uint64 `json:"roleId"       description:"角色ID"`
	RedTrace    int    `json:"redTrace" description:"红轨"`
	YellowTrace int    `json:"yellowTrace" description:"黄轨"`
	BlueTrace   int    `json:"blueTrace" description:"蓝轨"`
}

// TraceReduceReq 删除轨迹请求参数
type TraceReduceReq struct {
	g.Meta      `path:"/trace/Reduce" method:"delete" tags:"轨迹" summary:"删除轨迹"`
	RoleId      uint64 `json:"roleId"       description:"角色ID"`
	RedTrace    int    `json:"redTrace" description:"红轨"`
	YellowTrace int    `json:"yellowTrace" description:"黄轨"`
	BlueTrace   int    `json:"blueTrace" description:"蓝轨"`
}

// TraceReduceRes 删除轨迹响应数据
type TraceReduceRes struct {
}
