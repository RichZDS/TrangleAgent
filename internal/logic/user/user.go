package user

import (
	"context"
	v1 "leke/api/user/v1"
	"leke/internal/dao"
	"leke/internal/model"
	"leke/internal/service"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 创建数据库查询模型
	m := dao.Users.Ctx(ctx)

	// 按照账号搜索
	if req.Account != "" {
		m = m.WhereLike(dao.Users.Columns().Account, "%"+req.Account+"%")
	}

	// 按照昵称搜索
	if req.Nickname != "" {
		m = m.WhereLike(dao.Users.Columns().Nickname, "%"+req.Nickname+"%")
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var users []*model.User
	err = m.Scan(&users)
	if err != nil {
		return nil, err
	}

	// 转换为视图参数
	userViewList := make([]*model.UserViewParams, 0, len(users))
	for _, user := range users {
		userViewList = append(userViewList, &model.UserViewParams{
			Id:         user.Id,
			Account:    user.Account,
			Nickname:   user.Nickname,
			Gender:     user.Gender,
			BirthDate:  user.BirthDate,
			UserType:   user.UserType,
			VipStartAt: user.VipStartAt,
			VipEndAt:   user.VipEndAt,
			CreatedAt:  user.CreatedAt,
		})
	}

	res = &v1.UserListRes{
		Users: userViewList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

func (s *sUser) UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error) {
	// 创建数据库查询模型
	m := dao.Users.Ctx(ctx)

	// 根据账号查询
	if req.Account != "" {
		m = m.Where(dao.Users.Columns().Account, req.Account)
	}

	res = &v1.UserViewRes{}
	err = m.Scan(&res.UserViewParams)
	if err != nil {
		return nil, err
	}

	return
}

func (s *sUser) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	// 创建数据库查询模型
	m := dao.Users.Ctx(ctx)

	// 根据账号更新用户信息
	_, err = m.Data(req).Where(dao.Users.Columns().Account, req.Account).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.UserUpdateRes{}
	// 获取更新后的用户信息
	err = m.Where(dao.Users.Columns().Account, req.Account).Scan(&res.UserViewParams)
	if err != nil {
		return nil, err
	}

	return
}

func (s *sUser) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	// 根据账号删除用户
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Account, req.Account).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.UserDeleteRes{}
	return
}
