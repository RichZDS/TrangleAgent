package model

import "github.com/gogf/gf/v2/os/gtime"

type ClassListItem struct {
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
