// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GroupClassEnrollments is the golang structure for table group_class_enrollments.
type GroupClassEnrollments struct {
	Id          uint64      `json:"id"          orm:"id"            description:"报名ID"`                                    // 报名ID
	ClassId     uint64      `json:"classId"     orm:"class_id"      description:"团课ID"`                                    // 团课ID
	UserId      uint64      `json:"userId"      orm:"user_id"       description:"用户ID"`                                    // 用户ID
	Status      string      `json:"status"      orm:"status"        description:"状态：booked已报名，checked_in已签到，cancelled已取消"` // 状态：booked已报名，checked_in已签到，cancelled已取消
	EnrolledAt  *gtime.Time `json:"enrolledAt"  orm:"enrolled_at"   description:"报名时间"`                                    // 报名时间
	CheckedInAt *gtime.Time `json:"checkedInAt" orm:"checked_in_at" description:"签到时间"`                                    // 签到时间
}
