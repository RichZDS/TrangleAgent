package user

import (
	"context"
	v1 "leke/api/user/v1"
	"leke/internal/dao"
	"leke/internal/model"
	"leke/internal/model/entity"
	"leke/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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
	err = m.Where(dao.Users.Columns().Account, req.Account).Scan(&res)
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

// RoleCreate 创建角色
func (s *sUser) RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	// 插入角色数据
	data := &entity.RoleCards{
		UserId:       req.UserId,
		DepartmentId: req.DepartmentId,
		AgentName:    req.AgentName,
		CodeName:     req.CodeName,
		Gender:       req.Gender,
		ArcAbnormal:  req.ArcAbnormal,
		ArcReality:   req.ArcReality,
		ArcPosition:  req.ArcPosition,
	}

	id, err := dao.RoleCards.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	res = &v1.RoleCreateRes{
		Id: uint64(id),
	}
	return
}

// RoleUpdate 更新角色
func (s *sUser) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	// 更新角色数据
	data := &entity.RoleCards{
		DepartmentId: req.DepartmentId,
		AgentName:    req.AgentName,
		CodeName:     req.CodeName,
		Gender:       req.Gender,
		ArcAbnormal:  req.ArcAbnormal,
		ArcReality:   req.ArcReality,
		ArcPosition:  req.ArcPosition,
		Commendation: req.Commendation,
		Reprimand:    req.Reprimand,
		BlueTrack:    req.BlueTrack,
		YellowTrack:  req.YellowTrack,
		RedTrack:     req.RedTrack,
	}

	_, err = dao.RoleCards.Ctx(ctx).Data(data).Where(dao.RoleCards.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.RoleUpdateRes{
		Id: req.Id,
	}
	return
}

// RoleView 查看角色详情
func (s *sUser) RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error) {
	// 查询角色详情
	var role entity.RoleCards
	err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.Id).Scan(&role)
	if err != nil {
		return nil, err
	}

	res = &v1.RoleViewRes{
		RoleViewParams: v1.RoleViewParams{
			Id:             role.Id,
			UserId:         role.UserId,
			DepartmentId:   role.DepartmentId,
			Commendation:   role.Commendation,
			Reprimand:      role.Reprimand,
			BlueTrack:      role.BlueTrack,
			YellowTrack:    role.YellowTrack,
			RedTrack:       role.RedTrack,
			ArcAbnormal:    role.ArcAbnormal,
			ArcReality:     role.ArcReality,
			ArcPosition:    role.ArcPosition,
			AgentName:      role.AgentName,
			CodeName:       role.CodeName,
			Gender:         role.Gender,
			QaMeticulous:   role.QaMeticulous,
			QaDeception:    role.QaDeception,
			QaVigor:        role.QaVigor,
			QaEmpathy:      role.QaEmpathy,
			QaInitiative:   role.QaInitiative,
			QaResilience:   role.QaResilience,
			QaPresence:     role.QaPresence,
			QaProfessional: role.QaProfessional,
			QaDiscretion:   role.QaDiscretion,
		},
	}
	return
}

// RoleList 获取角色列表
func (s *sUser) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 根据用户ID查询
	if req.UserId != 0 {
		m = m.Where(dao.RoleCards.Columns().UserId, req.UserId)
	}

	// 根据部门ID查询
	if req.DepartmentId != 0 {
		m = m.Where(dao.RoleCards.Columns().DepartmentId, req.DepartmentId)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var roles []*entity.RoleCards
	err = m.Scan(&roles)
	if err != nil {
		return nil, err
	}

	// 转换为视图参数
	roleViewList := make([]*v1.RoleViewParams, 0, len(roles))
	for _, role := range roles {
		roleViewList = append(roleViewList, &v1.RoleViewParams{
			Id:             role.Id,
			UserId:         role.UserId,
			DepartmentId:   role.DepartmentId,
			Commendation:   role.Commendation,
			Reprimand:      role.Reprimand,
			BlueTrack:      role.BlueTrack,
			YellowTrack:    role.YellowTrack,
			RedTrack:       role.RedTrack,
			ArcAbnormal:    role.ArcAbnormal,
			ArcReality:     role.ArcReality,
			ArcPosition:    role.ArcPosition,
			AgentName:      role.AgentName,
			CodeName:       role.CodeName,
			Gender:         role.Gender,
			QaMeticulous:   role.QaMeticulous,
			QaDeception:    role.QaDeception,
			QaVigor:        role.QaVigor,
			QaEmpathy:      role.QaEmpathy,
			QaInitiative:   role.QaInitiative,
			QaResilience:   role.QaResilience,
			QaPresence:     role.QaPresence,
			QaProfessional: role.QaProfessional,
			QaDiscretion:   role.QaDiscretion,
		})
	}

	res = &v1.RoleListRes{
		List:       roleViewList,
		PageResult: req.PageResult, // 设置分页信息
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

// RoleDelete 删除角色
func (s *sUser) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	// 根据ID删除角色
	_, err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.RoleDeleteRes{
		Id: req.Id,
	}
	return
}

// RolePermissionCheck 权限查询
func (s *sUser) RolePermissionCheck(ctx context.Context, req *v1.RolePermissionCheckReq) (res *v1.RolePermissionCheckRes, err error) {
	// 查询角色信息
	var role entity.RoleCards
	err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.RoleId).Scan(&role)
	if err != nil {
		return nil, err
	}

	// 根据轨道类型获取对应的轨道值
	var trackValue uint
	switch req.TrackType {
	case "blue":
		trackValue = role.BlueTrack
	case "yellow":
		trackValue = role.YellowTrack
	case "red":
		trackValue = role.RedTrack
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "无效的轨道类型")
	}

	// 判断权限：如果角色的轨道值 >= 文件需要的值，则权限验证通过
	if trackValue >= req.FileValue {
		// 权限验证通过
		res = &v1.RolePermissionCheckRes{
			Code: 333,
			Mes:  "权限验证通过",
		}
		return
	}

	// 权限不足，返回错误
	return nil, gerror.New("权限不足")
}
