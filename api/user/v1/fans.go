package v1

import (
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FansListItem 粉丝列表项
type FansListItem struct {
	Id        uint64      `json:"id"         description:"自增主键"`
	UserId    uint64      `json:"userId"     description:"用户ID（被关注者，本人）"`
	FanId     uint64      `json:"fanId"      description:"粉丝用户ID"`
	Status    int         `json:"status"     description:"状态：1=粉丝 0=取关/无效"`
	Remark    string      `json:"remark"     description:"备注名（user_id 对 fan_id 的备注/分组）"`
	CreatedAt *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt"  description:"更新时间"`
}

// FansListReq 粉丝列表请求参数
type FansListReq struct {
	response.PageResult
	g.Meta `path:"/fans/list" method:"get" tags:"粉丝" summary:"粉丝列表"`
	UserId uint64 `json:"userId" v:"required#用户ID不能为空"`
	Status int    `json:"status" description:"状态：1=粉丝 0=取关/无效"`
}

// FansListRes 粉丝列表响应参数
type FansListRes struct {
	response.PageResult
	List []*FansListItem `json:"list"`
}

// FansViewReq 粉丝详情请求参数
type FansViewReq struct {
	g.Meta `path:"/fans/view" method:"get" tags:"粉丝" summary:"粉丝详情"`
	Id     uint64 `json:"id" v:"required#粉丝关系ID不能为空"`
}

// FansViewRes 粉丝详情响应参数
type FansViewRes struct {
	FansListItem
}

// FansUpdateReq 粉丝更新请求参数
type FansUpdateReq struct {
	g.Meta `path:"/fans/update" method:"put" tags:"粉丝" summary:"更新粉丝"`
	Id     uint64 `json:"id" v:"required#粉丝关系ID不能为空"`
	Status int    `json:"status"     description:"状态：1=粉丝 0=取关/无效"`
	Remark string `json:"remark"     description:"备注名（user_id 对 fan_id 的备注/分组）"`
}

// FansUpdateRes 粉丝更新响应参数
type FansUpdateRes struct {
	Id uint64 `json:"id"`
}

// FansDeleteReq 粉丝删除请求参数
type FansDeleteReq struct {
	g.Meta `path:"/fans/delete" method:"delete" tags:"粉丝" summary:"删除粉丝"`
	Id     uint64 `json:"id" v:"required#粉丝关系ID不能为空"`
}

// FansDeleteRes 粉丝删除响应参数
type FansDeleteRes struct {
}

// FansCreateReq 粉丝创建请求参数
type FansCreateReq struct {
	g.Meta `path:"/fans/create" method:"post" tags:"粉丝" summary:"创建粉丝"`
	UserId uint64 `json:"userId" v:"required#用户ID不能为空"`
	FanId  uint64 `json:"fanId" v:"required#粉丝用户ID不能为空"`
	Remark string `json:"remark" description:"备注名（user_id 对 fan_id 的备注/分组）"`
}

// FansCreateRes 粉丝创建响应参数
type FansCreateRes struct {
	Id uint64 `json:"id"`
}
