// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleCards is the golang structure for table role_cards.
type RoleCards struct {
	Id             uint64      `json:"id"             orm:"id"              description:"角色卡ID"`                                // 角色卡ID
	UserId         uint64      `json:"userId"         orm:"user_id"         description:"所属用户ID"`                               // 所属用户ID
	DepartmentId   uint64      `json:"departmentId"   orm:"department_id"   description:"所属部门ID"`                               // 所属部门ID
	Commendation   uint        `json:"commendation"   orm:"commendation"    description:"嘉奖次数"`                                 // 嘉奖次数
	Reprimand      uint        `json:"reprimand"      orm:"reprimand"       description:"申戒次数"`                                 // 申戒次数
	BlueTrack      uint        `json:"blueTrack"      orm:"blue_track"      description:"蓝轨（0-40）"`                             // 蓝轨（0-40）
	YellowTrack    uint        `json:"yellowTrack"    orm:"yellow_track"    description:"黄轨（0-40）"`                             // 黄轨（0-40）
	RedTrack       uint        `json:"redTrack"       orm:"red_track"       description:"红轨（0-40）"`                             // 红轨（0-40）
	ArcAbnormal    string      `json:"arcAbnormal"    orm:"arc_abnormal"    description:"ARC：异常"`                               // ARC：异常
	ArcReality     string      `json:"arcReality"     orm:"arc_reality"     description:"ARC：现实"`                               // ARC：现实
	ArcPosition    string      `json:"arcPosition"    orm:"arc_position"    description:"ARC：职位"`                               // ARC：职位
	AgentName      string      `json:"agentName"      orm:"agent_name"      description:"特工名字"`                                 // 特工名字
	CodeName       string      `json:"codeName"       orm:"code_name"       description:"代号"`                                   // 代号
	Gender         string      `json:"gender"         orm:"gender"          description:"性别"`                                   // 性别
	QaMeticulous   uint        `json:"qaMeticulous"   orm:"qa_meticulous"   description:"Meticulousness (0-100, QA)"`           // Meticulousness (0-100, QA)
	QaDeception    uint        `json:"qaDeception"    orm:"qa_deception"    description:"Deception (0-100, QA)"`                // Deception (0-100, QA)
	QaVigor        uint        `json:"qaVigor"        orm:"qa_vigor"        description:"Vigor / Drive (0-100, QA)"`            // Vigor / Drive (0-100, QA)
	QaEmpathy      uint        `json:"qaEmpathy"      orm:"qa_empathy"      description:"Empathy (0-100, QA)"`                  // Empathy (0-100, QA)
	QaInitiative   uint        `json:"qaInitiative"   orm:"qa_initiative"   description:"Initiative (0-100, QA)"`               // Initiative (0-100, QA)
	QaResilience   uint        `json:"qaResilience"   orm:"qa_resilience"   description:"Resilience / Persistence (0-100, QA)"` // Resilience / Persistence (0-100, QA)
	QaPresence     uint        `json:"qaPresence"     orm:"qa_presence"     description:"Presence / Charisma (0-100, QA)"`      // Presence / Charisma (0-100, QA)
	QaProfessional uint        `json:"qaProfessional" orm:"qa_professional" description:"Professionalism (0-100, QA)"`          // Professionalism (0-100, QA)
	QaDiscretion   uint        `json:"qaDiscretion"   orm:"qa_discretion"   description:"Discretion / Low profile (0-100, QA)"` // Discretion / Low profile (0-100, QA)
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`                                 // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`                                 // 更新时间
}
