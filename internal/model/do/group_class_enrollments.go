// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GroupClassEnrollments is the golang structure of table group_class_enrollments for DAO operations like Where/Data.
type GroupClassEnrollments struct {
	g.Meta      `orm:"table:group_class_enrollments, do:true"`
	Id          any         // 报名ID
	ClassId     any         // 团课ID
	UserId      any         // 用户ID
	Status      any         // 状态：booked已报名，checked_in已签到，cancelled已取消
	EnrolledAt  *gtime.Time // 报名时间
	CheckedInAt *gtime.Time // 签到时间
}
