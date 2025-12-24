package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"leke/internal/model/response"
)

type RoleCreateReq struct {
	g.Meta       `path:"/role/create" method:"post" tags:"角色" summary:"创建角色"`
	UserId       uint64 `json:"userId"         v:"required" description:"所属用户ID"`
	DepartmentId uint64 `json:"departmentId"   description:"所属部门ID"`
	AgentName    string `json:"agentName"      v:"required|length:1,50" description:"特工名字"`
	CodeName     string `json:"codeName"       v:"required|length:1,50" description:"代号"`
	Gender       string `json:"gender"         v:"required|in:male,female,other" description:"性别"`
	ArcAbnormal  string `json:"arcAbnormal"    description:"ARC：异常"`
	ArcReality   string `json:"arcReality"     description:"ARC：现实"`
	ArcPosition  string `json:"arcPosition"    description:"ARC：职位"`
}

type RoleCreateRes struct {
	Id uint64 `json:"id" description:"角色ID"`
}

type RoleUpdateReq struct {
	g.Meta       `path:"/role/update" method:"put" tags:"角色" summary:"更新角色"`
	Id           uint64 `json:"id"             v:"required" description:"角色ID"`
	DepartmentId uint64 `json:"departmentId"   description:"所属部门ID"`
	AgentName    string `json:"agentName"      v:"required|length:1,50" description:"特工名字"`
	CodeName     string `json:"codeName"       v:"required|length:1,50" description:"代号"`
	Gender       string `json:"gender"         v:"required|in:male,female,other" description:"性别"`
	ArcAbnormal  string `json:"arcAbnormal"    description:"ARC：异常"`
	ArcReality   string `json:"arcReality"     description:"ARC：现实"`
	ArcPosition  string `json:"arcPosition"    description:"ARC：职位"`
	Commendation uint   `json:"commendation"   description:"嘉奖次数"`
	Reprimand    uint   `json:"reprimand"      description:"申戒次数"`
	BlueTrack    uint   `json:"blueTrack"      description:"蓝轨（0-40）"`
	YellowTrack  uint   `json:"yellowTrack"    description:"黄轨（0-40）"`
	RedTrack     uint   `json:"redTrack"       description:"红轨（0-40）"`
}

type RoleUpdateRes struct {
	Id uint64 `json:"id" description:"角色ID"`
}

type RoleViewReq struct {
	g.Meta `path:"/role/view" method:"get" tags:"角色" summary:"查看角色"`
	Id     uint64 `json:"id" v:"required" description:"角色ID"`
}

type RoleViewRes struct {
	RoleViewParams `json:",inline"`
}

type RoleViewParams struct {
	Id             uint64 `json:"id"             description:"角色卡ID"`
	UserId         uint64 `json:"userId"         description:"所属用户ID"`
	DepartmentId   uint64 `json:"departmentId"   description:"所属部门ID"`
	Commendation   uint   `json:"commendation"   description:"嘉奖次数"`
	Reprimand      uint   `json:"reprimand"      description:"申戒次数"`
	BlueTrack      uint   `json:"blueTrack"      description:"蓝轨（0-40）"`
	YellowTrack    uint   `json:"yellowTrack"    description:"黄轨（0-40）"`
	RedTrack       uint   `json:"redTrack"       description:"红轨（0-40）"`
	ArcAbnormal    string `json:"arcAbnormal"    description:"ARC：异常"`
	ArcReality     string `json:"arcReality"     description:"ARC：现实"`
	ArcPosition    string `json:"arcPosition"    description:"ARC：职位"`
	AgentName      string `json:"agentName"      description:"特工名字"`
	CodeName       string `json:"codeName"       description:"代号"`
	Gender         string `json:"gender"         description:"性别"`
	QaMeticulous   uint   `json:"qaMeticulous"   description:"Meticulousness (0-100, QA)"`
	QaDeception    uint   `json:"qaDeception"    description:"Deception (0-100, QA)"`
	QaVigor        uint   `json:"qaVigor"        description:"Vigor / Drive (0-100, QA)"`
	QaEmpathy      uint   `json:"qaEmpathy"      description:"Empathy (0-100, QA)"`
	QaInitiative   uint   `json:"qaInitiative"   description:"Initiative (0-100, QA)"`
	QaResilience   uint   `json:"qaResilience"   description:"Resilience / Persistence (0-100, QA)"`
	QaPresence     uint   `json:"qaPresence"     description:"Presence / Charisma (0-100, QA)"`
	QaProfessional uint   `json:"qaProfessional" description:"Professionalism (0-100, QA)"`
	QaDiscretion   uint   `json:"qaDiscretion"   description:"Discretion / Low profile (0-100, QA)"`
}

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"角色列表"`
	response.PageResult
	UserId       uint64 `json:"userId"       description:"用户ID"`
	DepartmentId uint64 `json:"departmentId" description:"部门ID"`
}

type RoleListRes struct {
	response.PageResult
	List []*RoleViewParams `json:"list"  description:"角色列表"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色" summary:"删除角色"`
	Id     uint64 `json:"id" v:"required" description:"角色ID"`
}

type RoleDeleteRes struct {
	Id uint64 `json:"id" description:"角色ID"`
}

// RolePermissionCheckReq 权限查询请求
type RolePermissionCheckReq struct {
	g.Meta    `path:"/role/permission/check" method:"post" tags:"角色" summary:"检查角色权限"`
	RoleId    uint64 `json:"roleId"    v:"required" description:"角色ID"`
	FileValue uint   `json:"fileValue" v:"required|min:0|max:40" description:"文件需要的轨道值（0-40）"`
	TrackType string `json:"trackType" v:"required|in:blue,yellow,red" description:"轨道类型：blue(蓝轨),yellow(黄轨),red(红轨)"`
}

// RolePermissionCheckRes 权限查询响应
type RolePermissionCheckRes struct {
	Code int    `json:"code" description:"响应码：333表示成功"`
	Mes  string `json:"mes"  description:"响应消息"`
}
