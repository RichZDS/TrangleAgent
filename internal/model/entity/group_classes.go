// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GroupClasses is the golang structure for table group_classes.
type GroupClasses struct {
	Id          uint64      `json:"id"          orm:"id"           description:"团课ID"`               // 团课ID
	Title       string      `json:"title"       orm:"title"        description:"课程名称，如：燃脂搏击、瑜伽"`     // 课程名称，如：燃脂搏击、瑜伽
	Description string      `json:"description" orm:"description"  description:"课程介绍"`               // 课程介绍
	CoachId     uint64      `json:"coachId"     orm:"coach_id"     description:"授课教练ID"`             // 授课教练ID
	Location    string      `json:"location"    orm:"location"     description:"上课地点/门店/教室"`         // 上课地点/门店/教室
	StartTime   *gtime.Time `json:"startTime"   orm:"start_time"   description:"开始时间"`               // 开始时间
	EndTime     *gtime.Time `json:"endTime"     orm:"end_time"     description:"结束时间"`               // 结束时间
	MaxCapacity int         `json:"maxCapacity" orm:"max_capacity" description:"最大人数"`               // 最大人数
	Price       float64     `json:"price"       orm:"price"        description:"价格（0表示免费或已包含在会员卡内）"` // 价格（0表示免费或已包含在会员卡内）
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`               // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`               // 更新时间
}
