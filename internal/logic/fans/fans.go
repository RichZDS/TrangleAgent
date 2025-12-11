package fans

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	v1 "leke/api/user/v1"
	"leke/internal/dao"
	"leke/internal/service"
)

type sFans struct{}

func New() *sFans {
	return &sFans{}
}

func init() {
	service.RegisterFans(New())
}

// FansList 获取粉丝列表
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.FansListReq: 粉丝列表请求参数
//
// 返回:
//   - res *v1.FansListRes: 粉丝列表响应结果
//   - err error: 错误信息
func (s *sFans) FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error) {
	// 创建数据库查询模型
	m := dao.Fans.Ctx(ctx)

	// 根据用户ID查询
	m = m.Where(dao.Fans.Columns().UserId, req.UserId)

	// 根据状态查询
	if req.Status != 0 {
		m = m.Where(dao.Fans.Columns().Status, req.Status)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var fansList []*v1.FansListItem
	err = m.Scan(&fansList)
	if err != nil {
		return nil, err
	}

	res = &v1.FansListRes{
		List: fansList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

// FansView 获取粉丝详情
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.FansViewReq: 粉丝详情请求参数
//
// 返回:
//   - res *v1.FansViewRes: 粉丝详情响应结果
//   - err error: 错误信息
func (s *sFans) FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error) {
	// 创建数据库查询模型
	m := dao.Fans.Ctx(ctx)

	// 根据ID查询
	m = m.Where(dao.Fans.Columns().Id, req.Id)

	res = &v1.FansViewRes{}
	err = m.Scan(&res.FansListItem)
	if err != nil {
		return nil, err
	}

	return
}

// FansUpdate 更新粉丝信息
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.FansUpdateReq: 粉丝更新请求参数
//
// 返回:
//   - res *v1.FansUpdateRes: 粉丝更新响应结果
//   - err error: 错误信息
func (s *sFans) FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error) {
	// 创建数据库查询模型
	m := dao.Fans.Ctx(ctx)

	// 更新粉丝信息
	_, err = m.Data(req).Where(dao.Fans.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.FansUpdateRes{
		Id: req.Id,
	}

	return
}

// FansDelete 删除粉丝关系
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.FansDeleteReq: 粉丝删除请求参数
//
// 返回:
//   - res *v1.FansDeleteRes: 粉丝删除响应结果
//   - err error: 错误信息
func (s *sFans) FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error) {
	// 根据ID删除粉丝关系
	_, err = dao.Fans.Ctx(ctx).Where(dao.Fans.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.FansDeleteRes{}
	return
}

// FansCreate 创建粉丝关系
//
// 参数:
//   - ctx context.Context: 上下文信息
//   - req *v1.FansCreateReq: 粉丝创建请求参数
//
// 返回:
//   - res *v1.FansCreateRes: 粉丝创建响应结果
//   - err error: 错误信息
func (s *sFans) FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error) {
	// 创建数据库查询模型
	m := dao.Fans.Ctx(ctx)

	// 准备数据
	data := map[string]interface{}{
		"user_id":    req.UserId,
		"fan_id":     req.FanId,
		"status":     1, // 默认状态为关注中
		"remark":     req.Remark,
		"created_at": gtime.Now(),
		"updated_at": gtime.Now(),
	}

	// 插入数据
	result, err := m.Data(data).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	res = &v1.FansCreateRes{
		Id: uint64(lastInsertId),
	}

	return
}
