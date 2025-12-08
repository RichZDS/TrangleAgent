// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ContainmentRepo is the golang structure of table containment_repo for DAO operations like Where/Data.
type ContainmentRepo struct {
	g.Meta      `orm:"table:containment_repo, do:true"`
	Id          any // auto-increment primary key
	TerminalId  any // terminal (散逸端)
	Abnormal    any // number of contained anomalies (收容异常的数量)
	AnomalyName any // name of the anomaly (异常体的名字)
	AgentName   any // agent (特工)
	RepoName    any // containment repository name or code (收容库)
}
