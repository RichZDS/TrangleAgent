// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:users, do:true"`
	Id           any         // 用户ID
	Account      any         // 账号
	Password     any         // 密码哈希
	Nickname     any         // 昵称
	Gender       any         // 性别：0未知 1男 2女
	BirthDate    *gtime.Time // 生日
	UserType     any         // 用户类型：normal普通用户，vip为VIP用户
	VipStartAt   *gtime.Time // VIP开始时间
	VipEndAt     *gtime.Time // VIP结束时间
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	RealityRole  any         // 现实身份/角色
	AbnormalRole any         // 异常身份/角色
	JobTitle     any         // 职位
	Commendation any         //
	Admonition   any         //
	Email        any         //
	RedTrace     any         // 红轨
	YellowTrace  any         // 黄轨
	BlueTrace    any         // 蓝轨
}
