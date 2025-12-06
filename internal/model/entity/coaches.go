// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Coaches is the golang structure for table coaches.
type Coaches struct {
	Id        uint64      `json:"id"        orm:"id"         description:"教练ID"`                // 教练ID
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"对应的用户ID（如果教练也要登录的话）"` // 对应的用户ID（如果教练也要登录的话）
	Name      string      `json:"name"      orm:"name"       description:"教练姓名"`                // 教练姓名
	Phone     string      `json:"phone"     orm:"phone"      description:"教练电话"`                // 教练电话
	Specialty string      `json:"specialty" orm:"specialty"  description:"擅长方向，比如瑜伽、力量、减脂"`     // 擅长方向，比如瑜伽、力量、减脂
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                // 更新时间
}
