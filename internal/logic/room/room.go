package room

import (
	"context"
	v1 "leke/api/room/v1"
	"leke/internal/dao"
	"leke/internal/model"
	"leke/internal/model/entity"
	"leke/internal/service"
)

type sRoom struct{}

func New() *sRoom {
	return &sRoom{}
}

func init() {
	service.RegisterRoom(New())
}

func (s *sRoom) RoomList(ctx context.Context, req *v1.RoomListReq) (res *v1.RoomListRes, err error) {
	// 构建查询条件
	m := dao.TrpgRoom.Ctx(ctx)

	if req.RoomCode != "" {
		m = m.Where(dao.TrpgRoom.Columns().RoomCode, req.RoomCode)
	}
	if req.RoomName != "" {
		m = m.Where(dao.TrpgRoom.Columns().RoomName, req.RoomName)
	}
	if req.HostId != 0 {
		m = m.Where(dao.TrpgRoom.Columns().HostId, req.HostId)
	}
	if req.Status != 0 {
		m = m.Where(dao.TrpgRoom.Columns().Status, req.Status)
	}
	if req.SystemName != "" {
		m = m.Where(dao.TrpgRoom.Columns().SystemName, req.SystemName)
	}

	// 分页处理
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 查询总数
	totalCount, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 查询数据列表
	var rooms []*entity.TrpgRoom
	err = m.Page(req.Page, req.PageSize).Scan(&rooms)
	if err != nil {
		return nil, err
	}

	// 组装返回结果
	res = &v1.RoomListRes{
		PageResult: req.PageResult,
		Rooms:      make([]*model.RoomParams, len(rooms)),
	}

	// 转换数据格式
	for i, room := range rooms {
		res.Rooms[i] = &model.RoomParams{
			Id:             room.Id,
			RoomId:         room.RoomId,
			RoomCode:       room.RoomCode,
			RoomName:       room.RoomName,
			HostId:         room.HostId,
			HostNickname:   room.HostNickname,
			MaxPlayers:     room.MaxPlayers,
			CurrentPlayers: room.CurrentPlayers,
			HasPassword:    room.HasPassword,
			IsPrivate:      room.IsPrivate,
			Status:         room.Status,
			SystemName:     room.SystemName,
			ScenarioName:   room.ScenarioName,
			Description:    room.Description,
			CreatedAt:      room.CreatedAt,
			StartedAt:      room.StartedAt,
			EndedAt:        room.EndedAt,
			UpdatedAt:      room.UpdatedAt,
		}
	}

	// 设置分页信息
	res.PageResult.Total = totalCount
	res.PageResult.Page = req.Page
	res.PageResult.PageSize = req.PageSize

	return res, nil
}

func (s *sRoom) RoomView(ctx context.Context, req *v1.RoomViewReq) (res *v1.RoomViewRes, err error) {
	// 构建查询条件
	m := dao.TrpgRoom.Ctx(ctx)

	if req.Id != 0 {
		m = m.Where(dao.TrpgRoom.Columns().Id, req.Id)
	}
	if req.RoomCode != "" {
		m = m.Where(dao.TrpgRoom.Columns().RoomCode, req.RoomCode)
	}

	// 查询数据
	var room *entity.TrpgRoom
	err = m.Scan(&room)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil
	}

	// 组装返回结果
	res = &v1.RoomViewRes{
		RoomView: model.RoomView{
			Id:             room.Id,
			RoomId:         room.RoomId,
			RoomCode:       room.RoomCode,
			RoomName:       room.RoomName,
			HostId:         room.HostId,
			HostNickname:   room.HostNickname,
			MaxPlayers:     room.MaxPlayers,
			CurrentPlayers: room.CurrentPlayers,
			HasPassword:    room.HasPassword,
			IsPrivate:      room.IsPrivate,
			Status:         room.Status,
			SystemName:     room.SystemName,
			ScenarioName:   room.ScenarioName,
			Description:    room.Description,
			CreatedAt:      room.CreatedAt,
			StartedAt:      room.StartedAt,
			EndedAt:        room.EndedAt,
		},
	}

	return res, nil
}

func (s *sRoom) RoomUpdate(ctx context.Context, req *v1.RoomUpdateReq) (res *v1.RoomUpdateRes, err error) {
	// 根据房间号更新房间信息
	data := model.RoomParams{
		RoomName:       req.RoomName,
		HostId:         req.HostId,
		HostNickname:   req.HostNickname,
		MaxPlayers:     req.MaxPlayers,
		CurrentPlayers: req.CurrentPlayers,
		HasPassword:    req.HasPassword,
		IsPrivate:      req.IsPrivate,
		Status:         req.Status,
		SystemName:     req.SystemName,
		ScenarioName:   req.ScenarioName,
		Description:    req.Description,
	}

	_, err = dao.TrpgRoom.Ctx(ctx).Data(data).Where(dao.TrpgRoom.Columns().RoomCode, req.RoomCode).Update()
	if err != nil {
		return nil, err
	}

	// 获取更新后的房间ID
	var room *entity.TrpgRoom
	err = dao.TrpgRoom.Ctx(ctx).Where(dao.TrpgRoom.Columns().RoomCode, req.RoomCode).Scan(&room)
	if err != nil {
		return nil, err
	}

	res = &v1.RoomUpdateRes{
		Id: room.Id,
	}

	return res, nil
}

func (s *sRoom) RoomDelete(ctx context.Context, req *v1.RoomDeleteReq) (res *v1.RoomDeleteRes, err error) {
	// 根据房间号删除房间
	_, err = dao.TrpgRoom.Ctx(ctx).Where(dao.TrpgRoom.Columns().RoomCode, req.RoomCode).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.RoomDeleteRes{}
	return res, nil
}

func (s *sRoom) RoomCreate(ctx context.Context, req *v1.RoomCreateReq) (res *v1.RoomCreateRes, err error) {
	// 准备房间数据
	room := entity.TrpgRoom{
		RoomCode:     req.RoomCode,
		RoomName:     req.RoomName,
		HostId:       req.HostId,
		MaxPlayers:   req.MaxPlayers,
		HasPassword:  req.HasPassword,
		IsPrivate:    req.IsPrivate,
		SystemName:   req.SystemName,
		ScenarioName: req.ScenarioName,
		Description:  req.Description,
		Status:       0, // 默认状态为未开始
	}

	// 插入数据库
	result, err := dao.TrpgRoom.Ctx(ctx).Data(room).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入的ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 返回结果
	res = &v1.RoomCreateRes{
		Id:       uint64(id),
		RoomCode: req.RoomCode,
	}

	return res, nil
}
