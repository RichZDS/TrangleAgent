// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkoutLogs is the golang structure for table workout_logs.
type WorkoutLogs struct {
	Id              uint64      `json:"id"              orm:"id"               description:"健身记录ID"`             // 健身记录ID
	UserId          uint64      `json:"userId"          orm:"user_id"          description:"关联的用户ID"`            // 关联的用户ID
	WorkoutDate     *gtime.Time `json:"workoutDate"     orm:"workout_date"     description:"健身日期"`               // 健身日期
	StartTime       *gtime.Time `json:"startTime"       orm:"start_time"       description:"开始时间"`               // 开始时间
	EndTime         *gtime.Time `json:"endTime"         orm:"end_time"         description:"结束时间"`               // 结束时间
	DurationMinutes int         `json:"durationMinutes" orm:"duration_minutes" description:"时长（分钟）"`             // 时长（分钟）
	Calories        int         `json:"calories"        orm:"calories"         description:"估算消耗的卡路里"`           // 估算消耗的卡路里
	Description     string      `json:"description"     orm:"description"      description:"本次训练的备注，比如训练内容、部位等"` // 本次训练的备注，比如训练内容、部位等
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"记录创建时间"`             // 记录创建时间
}
