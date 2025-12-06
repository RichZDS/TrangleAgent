// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GroupClasses is the golang structure of table group_classes for DAO operations like Where/Data.
type GroupClasses struct {
	g.Meta      `orm:"table:group_classes, do:true"`
	Id          any         // 团课ID
	Title       any         // 课程名称，如：燃脂搏击、瑜伽
	Description any         // 课程介绍
	CoachId     any         // 授课教练ID
	Location    any         // 上课地点/门店/教室
	StartTime   *gtime.Time // 开始时间
	EndTime     *gtime.Time // 结束时间
	MaxCapacity any         // 最大人数
	Price       any         // 价格（0表示免费或已包含在会员卡内）
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
