package v1

import (
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SubscribeListItem 关注列表项
type SubscribeListItem struct {
	Id        uint64      `json:"id"         description:"自增主键"`
	UserId    uint64      `json:"userId"     description:"用户ID（关注者，本人）"`
	FollowId  uint64      `json:"followId"   description:"被关注的用户ID"`
	Status    int         `json:"status"     description:"状态：1=关注中 0=取消关注"`
	Remark    string      `json:"remark"     description:"备注名（user_id 对 follow_id 的备注）"`
	CreatedAt *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt"  description:"更新时间"`
}

// SubscribeListReq 关注列表请求参数
type SubscribeListReq struct {
	response.PageResult
	g.Meta   `path:"/subscribe/list" method:"get" tags:"关注" summary:"关注列表"`
	UserId   uint64 `json:"userId" v:"required#用户ID不能为空"`
	Status   int    `json:"status" description:"状态：1=关注中 0=取消关注"`
	FollowId uint64 `json:"followId" description:"被关注的用户ID"`
}

// SubscribeListRes 关注列表响应参数
type SubscribeListRes struct {
	response.PageResult
	List []*SubscribeListItem `json:"list"`
}

// SubscribeViewReq 关注详情请求参数
type SubscribeViewReq struct {
	g.Meta `path:"/subscribe/view" method:"get" tags:"关注" summary:"关注详情"`
	Id     uint64 `json:"id" v:"required#关注关系ID不能为空"`
}

// SubscribeViewRes 关注详情响应参数
type SubscribeViewRes struct {
	SubscribeListItem
}

// SubscribeUpdateReq 关注更新请求参数
type SubscribeUpdateReq struct {
	g.Meta `path:"/subscribe/update" method:"put" tags:"关注" summary:"更新关注"`
	Id     uint64 `json:"id" v:"required#关注关系ID不能为空"`
	Status int    `json:"status"     description:"状态：1=关注中 0=取消关注"`
	Remark string `json:"remark"     description:"备注名（user_id 对 follow_id 的备注）"`
}

// SubscribeUpdateRes 关注更新响应参数
type SubscribeUpdateRes struct {
	Id uint64 `json:"id"`
}

// SubscribeDeleteReq 关注删除请求参数
type SubscribeDeleteReq struct {
	g.Meta `path:"/subscribe/delete" method:"delete" tags:"关注" summary:"删除关注"`
	Id     uint64 `json:"id" v:"required#关注关系ID不能为空"`
}

// SubscribeDeleteRes 关注删除响应参数
type SubscribeDeleteRes struct {
}

// SubscribeCreateReq 关注创建请求参数
type SubscribeCreateReq struct {
	g.Meta   `path:"/subscribe/create" method:"post" tags:"关注" summary:"创建关注"`
	UserId   uint64 `json:"userId" v:"required#用户ID不能为空"`
	FollowId uint64 `json:"followId" v:"required#被关注用户ID不能为空"`
	Remark   string `json:"remark" description:"备注名（user_id 对 follow_id 的备注）"`
}

// SubscribeCreateRes 关注创建响应参数
type SubscribeCreateRes struct {
	Id uint64 `json:"id"`
}
