package v1

import (
	"leke/internal/model"
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	response.PageResult
	g.Meta   `path:"/user/list" method:"get" tags:"User" summary:"User List"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
}

type UserListRes struct {
	response.PageResult
	Users []*model.UserViewParams
}

type UserViewReq struct {
	g.Meta   `path:"/user/view" method:"get" tags:"User" summary:"User View"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
}
type UserViewRes struct {
	model.UserViewParams
}
type UserUpdateReq struct {
	g.Meta  `path:"/user/update" method:"put" tags:"User" summary:"User Update"`
	Account string `json:"account"    orm:"account"      description:"账号"`
}
type UserUpdateRes struct {
	model.UserViewParams
}
type UserDeleteReq struct {
	g.Meta  `path:"/user/delete" method:"delete" tags:"User" summary:"User Delete"`
	Account string `json:"account"    orm:"account"      description:"账号"`
}
type UserDeleteRes struct {
}
