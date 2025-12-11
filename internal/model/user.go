package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id         uint64      `json:"id"         orm:"id"           description:"用户ID"`                      // 用户ID
	Account    string      `json:"account"    orm:"account"      description:"账号"`                        // 账号
	Password   string      `json:"password"   orm:"password"     description:"密码哈希"`                      // 密码哈希
	Nickname   string      `json:"nickname"   orm:"nickname"     description:"昵称"`                        // 昵称
	Gender     int         `json:"gender"     orm:"gender"       description:"性别：0未知 1男 2女"`              // 性别：0未知 1男 2女
	BirthDate  *gtime.Time `json:"birthDate"  orm:"birth_date"   description:"生日"`                        // 生日
	UserType   string      `json:"userType"   orm:"user_type"    description:"用户类型：normal普通用户，vip为VIP用户"` // 用户类型：normal普通用户，vip为VIP用户
	VipStartAt *gtime.Time `json:"vipStartAt" orm:"vip_start_at" description:"VIP开始时间"`                   // VIP开始时间
	VipEndAt   *gtime.Time `json:"vipEndAt"   orm:"vip_end_at"   description:"VIP结束时间"`                   // VIP结束时间
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`                      // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"更新时间"`                      // 更新时间
}

type UserViewParams struct {
	Id           uint64      `json:"id"         orm:"id"           description:"用户ID"`                      // 用户ID
	Account      string      `json:"account"    orm:"account"      description:"账号"`                        // 账号
	Nickname     string      `json:"nickname"   orm:"nickname"     description:"昵称"`                        // 昵称
	Gender       int         `json:"gender"     orm:"gender"       description:"性别：0未知 1男 2女"`              // 性别：0未知 1男 2女
	BirthDate    *gtime.Time `json:"birthDate"  orm:"birth_date"   description:"生日"`                        // 生日
	UserType     string      `json:"userType"   orm:"user_type"    description:"用户类型：normal普通用户，vip为VIP用户"` // 用户类型：normal普通用户，vip为VIP用户
	VipStartAt   *gtime.Time `json:"vipStartAt" orm:"vip_start_at" description:"VIP开始时间"`                   // VIP开始时间
	VipEndAt     *gtime.Time `json:"vipEndAt"   orm:"vip_end_at"   description:"VIP结束时间"`                   // VIP结束时间
	CreatedAt    *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`                      // 创建时间
	RealityRole  string      `json:"realityRole"  orm:"reality_role"   description:"现实身份/角色"`               // 现实身份/角色
	AbnormalRole string      `json:"abnormalRole" orm:"abnormal_role"  description:"异常身份/角色"`               // 异常身份/角色
	JobTitle     string      `json:"jobTitle"     orm:"job_title"      description:"职位"`                    // 职位
}
