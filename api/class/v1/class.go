package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"leke/internal/model"
	"leke/internal/model/response"
)

// 课程列表请求参数
type ClassListReq struct {
	g.Meta `path:"/class/list" method:"get" tags:"Class" summary:"Class List"`
	Title  string `json:"title"   description:"课程名称"`
	Id     uint64 `json:"id"      description:"课程ID"`
	response.PageResult
}

// 课程列表返回结果
type ClassListRes struct {
	Classes []*model.ClassListItem `json:"classes"`
}

// 课程详情请求参数
type ClassViewReq struct {
	g.Meta `path:"/class/view" method:"get" tags:"Class" summary:"Class View"`
	Id     uint64 `json:"id" v:"required" description:"课程ID"`
}

// 课程详情返回结果
type ClassViewRes struct {
	Class *ClassInfo `json:"class"`
}

type ClassInfo struct {
	Id          uint64      `json:"id"           description:"团课ID"`
	Title       string      `json:"title"        description:"课程名称，如：燃脂搏击、瑜伽"`
	Description string      `json:"description"  description:"课程介绍"`
	CoachId     uint64      `json:"coachId"      description:"授课教练ID"`
	Location    string      `json:"location"     description:"上课地点/门店/教室"`
	StartTime   *gtime.Time `json:"startTime"    description:"开始时间"`
	EndTime     *gtime.Time `json:"endTime"      description:"结束时间"`
	MaxCapacity int         `json:"maxCapacity"  description:"最大人数"`
	Price       float64     `json:"price"        description:"价格（0表示免费或已包含在会员卡内）"`
	CreatedAt   *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"    description:"更新时间"`
}

// 课程更新请求参数
type ClassUpdateReq struct {
	g.Meta      `path:"/class/update" method:"put" tags:"Class" summary:"Class Update"`
	Id          uint64      `json:"id"           description:"课程ID"`
	Title       string      `json:"title"        description:"课程名称，如：燃脂搏击、瑜伽"`
	Description string      `json:"description"  description:"课程介绍"`
	CoachId     uint64      `json:"coachId"      description:"授课教练ID"`
	Location    string      `json:"location"     description:"上课地点/门店/教室"`
	StartTime   *gtime.Time `json:"startTime"    description:"开始时间"`
	EndTime     *gtime.Time `json:"endTime"      description:"结束时间"`
	MaxCapacity int         `json:"maxCapacity"  description:"最大人数"`
	Price       float64     `json:"price"        description:"价格（0表示免费或已包含在会员卡内）"`
	CreatedAt   *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"    description:"更新时间"`
}

// 课程更新返回结果
type ClassUpdateRes struct {
	Updated bool `json:"updated"`
}

// 课程删除请求参数
type ClassDeleteReq struct {
	g.Meta `path:"/class/delete" method:"delete" tags:"Class" summary:"Class Delete"`
	Id     uint64 `json:"id" v:"required" description:"课程ID"`
}

// 课程删除返回结果
type ClassDeleteRes struct {
	Deleted bool `json:"deleted"`
}
