package containment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "leke/api/containment/v1"
	"leke/internal/model/response"
	"leke/internal/service"
)

type sContainment struct{}

func New() *sContainment {
	return &sContainment{}
}

func init() {
	service.RegisterContainment(New())
}

func (s *sContainment) ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error) {
	m := g.Model("containment_repo")
	if req.AgentName != "" {
		m.Where("agent_name", req.AgentName)
	}
	if req.AnomalyName != "" {
		m.Where("anomaly_name", req.AnomalyName)
	}
	if req.RepoName != "" {
		m.Where("repo_name", req.RepoName)
	}

	// 创建返回结果
	res = &v1.ContainmentRepoListRes{
		PageResult: response.PageResult{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}

	// 查询总数
	res.Total, err = m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	m = m.Page(req.Page, req.PageSize)

	// 查询数据列表
	err = m.Scan(&res.List)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *sContainment) ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error) {
	// 创建返回结果
	res = &v1.ContainmentRepoViewRes{}

	// 从数据库获取指定ID的数据
	err = g.Model("containment_repo").Where("id", req.Id).Scan(&res.ContainmentRepoInfo)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *sContainment) ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error) {
	// 创建返回结果
	res = &v1.ContainmentRepoUpdateRes{}

	// 准备数据
	data := g.Map{
		"terminal_id":  req.TerminalId,
		"anomaly_name": req.AnomalyName,
		"agent_name":   req.AgentName,
		"repo_name":    req.RepoName,
	}

	// 根据ID判断是更新还是新增
	if req.Id > 0 {
		// 更新操作
		_, err = g.Model("containment_repo").Data(data).Where("id", req.Id).Update()
		res.Id = req.Id
	} else {
		// 新增操作
		result, err := g.Model("containment_repo").Data(data).Insert()
		if err != nil {
			return nil, err
		}

		// 获取插入记录的ID
		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}
		res.Id = uint64(id)
	}

	return res, err
}
func (s *sContainment) ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error) {
	// 创建返回结果
	res = &v1.ContainmentRepoDeleteRes{}

	// 删除指定ID的记录
	_, err = g.Model("containment_repo").Where("id", req.Id).Delete()
	if err != nil {
		return nil, err
	}

	return res, nil
}
