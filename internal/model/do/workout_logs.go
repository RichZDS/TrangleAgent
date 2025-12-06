// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkoutLogs is the golang structure of table workout_logs for DAO operations like Where/Data.
type WorkoutLogs struct {
	g.Meta          `orm:"table:workout_logs, do:true"`
	Id              any         // 健身记录ID
	UserId          any         // 关联的用户ID
	WorkoutDate     *gtime.Time // 健身日期
	StartTime       *gtime.Time // 开始时间
	EndTime         *gtime.Time // 结束时间
	DurationMinutes any         // 时长（分钟）
	Calories        any         // 估算消耗的卡路里
	Description     any         // 本次训练的备注，比如训练内容、部位等
	CreatedAt       *gtime.Time // 记录创建时间
}
