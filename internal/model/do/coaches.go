// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Coaches is the golang structure of table coaches for DAO operations like Where/Data.
type Coaches struct {
	g.Meta    `orm:"table:coaches, do:true"`
	Id        any         // 教练ID
	UserId    any         // 对应的用户ID（如果教练也要登录的话）
	Name      any         // 教练姓名
	Phone     any         // 教练电话
	Specialty any         // 擅长方向，比如瑜伽、力量、减脂
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
