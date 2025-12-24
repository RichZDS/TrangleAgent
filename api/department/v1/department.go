package v1

import (
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Department 三角机构分部信息
type Department struct {
	Id            uint64      `json:"id"           description:"自增主键"`
	UserId        uint64      `json:"userId"       description:"所属用户ID（对应 users.id）"`
	BranchName    string      `json:"branchName"   description:"分部名称"`
	TerminalCount int         `json:"terminalCount" description:"分部散逸端的数量"`
	Weather       string      `json:"weather"      description:"分部当前天气/气候描述"`
	ManagerName   string      `json:"managerName"  description:"分部经理名称"`
	Location      string      `json:"location"     description:"分部地址"`
	CreatedAt     *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"    description:"更新时间"`
}

// DepartmentListReq 部门列表请求参数
type DepartmentListReq struct {
	response.PageResult
	g.Meta      `path:"/department/list" method:"get" tags:"部门" summary:"部门列表"`
	BranchName  string `json:"branchName"   description:"分部名称"`
	ManagerName string `json:"managerName"  description:"分部经理名称"`
	UserId      uint64 `json:"userId"       description:"所属用户ID"`
}

// DepartmentListRes 部门列表响应参数
type DepartmentListRes struct {
	response.PageResult
	List []*Department `json:"list"`
}

// DepartmentViewReq 部门详情请求参数
type DepartmentViewReq struct {
	g.Meta `path:"/department/view" method:"get" tags:"部门" summary:"部门详情"`
	Id     uint64 `json:"id" v:"required#部门ID不能为空"`
}

// DepartmentViewRes 部门详情响应参数
type DepartmentViewRes struct {
	Department
}

// DepartmentCreateReq 部门创建请求参数
type DepartmentCreateReq struct {
	g.Meta        `path:"/department/create" method:"post" tags:"部门" summary:"创建部门"`
	UserId        uint64 `json:"userId"        v:"required#所属用户ID不能为空"`
	BranchName    string `json:"branchName"    v:"required#分部名称不能为空"`
	TerminalCount int    `json:"terminalCount" description:"分部散逸端的数量"`
	Weather       string `json:"weather"       description:"分部当前天气/气候描述"`
	ManagerName   string `json:"managerName"   description:"分部经理名称"`
	Location      string `json:"location"     description:"分部地址"`
}

// DepartmentCreateRes 部门创建响应参数
type DepartmentCreateRes struct {
	Id uint64 `json:"id"`
}

// DepartmentUpdateReq 部门更新请求参数
type DepartmentUpdateReq struct {
	g.Meta        `path:"/department/update" method:"put" tags:"部门" summary:"更新部门"`
	Id            uint64 `json:"id"            v:"required#部门ID不能为空"`
	UserId        uint64 `json:"userId"        v:"required#所属用户ID不能为空"`
	BranchName    string `json:"branchName"    v:"required#分部名称不能为空"`
	TerminalCount int    `json:"terminalCount" description:"分部散逸端的数量"`
	Location      string `json:"location"     description:"分部地址"`
	Weather       string `json:"weather"       description:"分部当前天气/气候描述"`
	ManagerName   string `json:"managerName"   description:"分部经理名称"`
}

// DepartmentUpdateRes 部门更新响应参数
type DepartmentUpdateRes struct {
	Id uint64 `json:"id"`
}

// DepartmentDeleteReq 部门删除请求参数
type DepartmentDeleteReq struct {
	g.Meta `path:"/department/delete" method:"delete" tags:"部门" summary:"删除部门"`
	Id     uint64 `json:"id" v:"required#部门ID不能为空"`
}

// DepartmentDeleteRes 部门删除响应参数
type DepartmentDeleteRes struct {
}
