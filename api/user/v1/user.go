package v1

import (
	"leke/internal/model"
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	response.PageResult
	g.Meta   `path:"/user/list" method:"get" tags:"用户" summary:"用户列表"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
}

type UserListRes struct {
	response.PageResult
	Users []*model.UserViewParams
}

type UserViewReq struct {
	g.Meta   `path:"/user/view" method:"get" tags:"用户" summary:"用户详情"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
	Id       uint64 `json:"id"         orm:"id"           description:"用户ID"`
}
type UserViewRes struct {
	model.UserViewParams
}
type UserUpdateReq struct {
	g.Meta  `path:"/user/update" method:"put" tags:"用户" summary:"更新用户"`
	Account string `json:"account"    orm:"account"      description:"账号"`
	model.User
}
type UserUpdateRes struct {
	Id uint64 `json:"id"         orm:"id"           description:"用户ID"`
}
type UserDeleteReq struct {
	g.Meta  `path:"/user/delete" method:"delete" tags:"用户" summary:"删除用户"`
	Account string `json:"account"    orm:"account"      description:"账号"`
}
type UserDeleteRes struct {
}
