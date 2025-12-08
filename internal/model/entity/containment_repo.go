// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ContainmentRepo is the golang structure for table containment_repo.
type ContainmentRepo struct {
	Id            uint64 `json:"id"            orm:"id"             description:"auto-increment primary key"`                // auto-increment primary key
	TerminalId    int    `json:"terminalId"    orm:"terminal_id"    description:"terminal (散逸端)"`                            // terminal (散逸端)
	AbnormalCount int    `json:"abnormalCount" orm:"abnormal_count" description:"number of contained anomalies (收容异常的数量)"`   // number of contained anomalies (收容异常的数量)
	AnomalyName   string `json:"anomalyName"   orm:"anomaly_name"   description:"name of the anomaly (异常体的名字)"`              // name of the anomaly (异常体的名字)
	AgentName     string `json:"agentName"     orm:"agent_name"     description:"agent (特工)"`                                // agent (特工)
	RepoName      string `json:"repoName"      orm:"repo_name"      description:"containment repository name or code (收容库)"` // containment repository name or code (收容库)
}
