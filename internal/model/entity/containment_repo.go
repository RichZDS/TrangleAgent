// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ContainmentRepo is the golang structure for table containment_repo.
type ContainmentRepo struct {
	Id          uint64 `json:"id"          orm:"id"           description:"auto-increment primary key"`                // auto-increment primary key
	AnomalyName string `json:"anomalyName" orm:"anomaly_name" description:"name of the anomaly (异常体的名字)"`              // name of the anomaly (异常体的名字)
	AgentName   string `json:"agentName"   orm:"agent_name"   description:"agent (特工)"`                                // agent (特工)
	RepoName    string `json:"repoName"    orm:"repo_name"    description:"containment repository name or code (收容库)"` // containment repository name or code (收容库)
	Department  int    `json:"department"  orm:"department"   description:"部门"`                                        // 部门
}
