package model

type ContainmentRepo struct {
	Abnormal    int    `json:"abnormal"     orm:"abnormal"     description:"number of contained anomalies (收容异常的数量)"`
	AgentName   string `json:"agentName"     orm:"agent_name"     description:"agent (特工)"`
	AnomalyName string `json:"anomalyName"   orm:"anomaly_name"   description:"name of the anomaly (异常体的名字)"`
	Id          uint64 `json:"id"            orm:"id"             description:"primary key(大于0为修改，其他为新增)"`
	RepoName    string `json:"repoName"      orm:"repo_name"      description:"containment repository name or code (收容库)"`
	TerminalId  int    `json:"terminalId"    orm:"terminal_id"    description:"terminal (散逸端)"`
}
