package class

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "leke/api/class/v1"
	"leke/internal/dao"
	"leke/internal/model"
	"leke/internal/service"
)

type sClass struct{}

func New() *sClass {
	return &sClass{}
}

// 关键！！在 init 里完成注册
func init() {
	service.RegisterClass(New())
}

func (s *sClass) ClassList(ctx context.Context, req *v1.ClassListReq) (res *v1.ClassListRes, err error) {
	// 创建数据库查询模型
	m := dao.GroupClasses.Ctx(ctx)

	// 按照课程名字搜索
	if req.Title != "" {
		m = m.WhereLike(dao.GroupClasses.Columns().Title, "%"+req.Title+"%")
	}

	// 如果指定了ID，则按ID查询
	if req.Id != 0 {
		m = m.Where(dao.GroupClasses.Columns().Id, req.Id)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	classes := make([]*model.ClassListItem, 0)
	err = m.Scan(&classes)
	if err != nil {
		return nil, err
	}

	res = &v1.ClassListRes{
		Classes: classes,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

func (s *sClass) ClassView(ctx context.Context, req *v1.ClassViewReq) (res *v1.ClassViewRes, err error) {
	res = &v1.ClassViewRes{}
	err = dao.GroupClasses.Ctx(ctx).Where(dao.GroupClasses.Columns().Id, req.Id).Scan(&res.Class)
	return
}
func (s *sClass) ClassUpdate(ctx context.Context, req *v1.ClassUpdateReq) (res *v1.ClassUpdateRes, err error) {
	res = &v1.ClassUpdateRes{}

	if req.Id > 0 {
		// 修改
		req.UpdatedAt = gtime.Now()
		req.CreatedAt = gtime.Now()

		_, err = dao.GroupClasses.Ctx(ctx).Where(dao.GroupClasses.Columns().Id, req.Id).Data(req).Update()
		if err != nil {
			return nil, err
		}
	} else {
		// 新增
		req.UpdatedAt = gtime.Now()
		_, err = dao.GroupClasses.Ctx(ctx).Data(req).Insert()
		if err != nil {
			return nil, err
		}
	}

	res.Updated = true
	return res, nil
}
func (s *sClass) ClassDelete(ctx context.Context, req *v1.ClassDeleteReq) (res *v1.ClassDeleteRes, err error) {
	_, err = dao.GroupClasses.Ctx(ctx).Where(dao.GroupClasses.Columns().Id, req.Id).Delete()
	return &v1.ClassDeleteRes{Deleted: true}, err
}
