package v1

import (
	"leke/internal/model"
	"leke/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// 单条数据结构
type ContainmentRepoInfo struct {
	Id          uint64 `json:"id"          orm:"id"           description:"primary key(大于0为修改，其他为新增)"`
	TerminalId  int    `json:"terminalId"  orm:"terminal_id"  description:"terminal (散逸端)"`
	Abnormal    string `json:"abnormal"    orm:"abnormal"     description:"number of contained anomalies (收容异常的数量)"`
	AnomalyName string `json:"anomalyName" orm:"anomaly_name" description:"name of the anomaly (异常体的名字)"`
	AgentName   string `json:"agentName"   orm:"agent_name"   description:"agent (特工)"`
	RepoName    string `json:"repoName"    orm:"repo_name"    description:"containment repository name or code (收容库)"`
}

// List
type ContainmentRepoListReq struct {
	response.PageResult
	g.Meta `path:"/containment/list" method:"get" tags:"收容库"  summary:"收容库列表"`

	// 查询条件（全部可选）
	TerminalId  int    `json:"terminalId"  orm:"terminal_id"  description:"terminal (散逸端)"`
	AnomalyName string `json:"anomalyName" orm:"anomaly_name" description:"anomaly name"`
	AgentName   string `json:"agentName"   orm:"agent_name"   description:"agent name"`
	RepoName    string `json:"repoName"    orm:"repo_name"    description:"containment repo name"`
}

type ContainmentRepoListRes struct {
	response.PageResult
	List []*model.ContainmentRepo `json:"list"`
}

// View
type ContainmentRepoViewReq struct {
	g.Meta `path:"/containment/view" method:"get" tags:"收容库" summary:"收容库详情"`

	Id uint64 `json:"id" orm:"id" description:"primary key"`
}

type ContainmentRepoViewRes struct {
	ContainmentRepoInfo
}

// Update（id>0 为修改，id<=0 或不传为新增）
type ContainmentRepoUpdateReq struct {
	g.Meta `path:"/containment/update" method:"post" tags:"收容库" summary:"创建或更新收容库"`

	Id          uint64 `json:"id"          orm:"id"           description:"primary key，大于0为修改，其他为新增"`
	TerminalId  int    `json:"terminalId"  orm:"terminal_id"  description:"terminal (散逸端)"`
	Abnormal    string `json:"abnormal"    orm:"abnormal"     description:"number of contained anomalies"`
	AnomalyName string `json:"anomalyName" orm:"anomaly_name" description:"anomaly name"`
	AgentName   string `json:"agentName"   orm:"agent_name"   description:"agent name"`
	RepoName    string `json:"repoName"    orm:"repo_name"    description:"containment repo name"`
}

type ContainmentRepoUpdateRes struct {
	Id uint64 `json:"id" orm:"id" description:"primary key"`
}

// Delete
type ContainmentRepoDeleteReq struct {
	g.Meta `path:"/containment/delete" method:"delete" tags:"收容库" summary:"收容库删除"`

	Id uint64 `json:"id" orm:"id" description:"primary key"`
}

type ContainmentRepoDeleteRes struct{}
